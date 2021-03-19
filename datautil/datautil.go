package datautil

import (
	"regexp"

	"github.com/scutrobotlab/asuwave/variable"
)

func MakeCmd(act uint8, v *variable.T) []byte {
	data := make([]byte, 16)
	data[0] = byte(v.Board)
	data[1] = act
	data[2] = byte(variable.TypeLen[v.Type])
	copy(data[3:7], variable.AnyToBytes(v.Addr))
	if act == ActModeWrite {
		copy(data[7:15], variable.SpecToBytes(v.Type, v.Data))
	}
	data[15] = '\n'
	return data
}

func FindValidPart(data []byte) (int, int) {
	n := len(data)
	if n < 20 {
		return 0, 0
	}

	l := 0
	r := 19

	for r < n && (data[l] != Board1 || data[l+1] != ActModeSubscribeReturn || data[r] != '\n') {
		l++
		r++
	}

	if r == n {
		return 0, 0
	}

	r = n - (n-r)%20
	return l, r + 1
}

func MakeChartPack(x *variable.ListChartT, y *variable.ListT, data []byte) {
	var chartData variable.ToChartT
	for i := 0; i < len(data)/20; i++ {
		chartData.Board = data[i*20]
		chartData.Tick = variable.BytesToUint32(data[i*20+15 : i*20+19])
		addr := variable.BytesToUint32(data[i*20+3 : i*20+7])

		for _, v := range y.Variables {
			if v.Addr == addr {
				chartData.Name = v.Name
				switch v.Type {
				case "uint8_t":
					chartData.Data = float64(variable.BytesToUint8(data[i*20+7 : i*20+15]))
				case "uint16_t":
					chartData.Data = float64(variable.BytesToUint16(data[i*20+7 : i*20+15]))
				case "uint32_t":
					chartData.Data = float64(variable.BytesToUint32(data[i*20+7 : i*20+15]))
				case "uint64_t":
					chartData.Data = float64(variable.BytesToUint64(data[i*20+7 : i*20+15]))
				case "int8_t":
					chartData.Data = float64(variable.BytesToInt8(data[i*20+7 : i*20+15]))
				case "int16_t":
					chartData.Data = float64(variable.BytesToInt16(data[i*20+7 : i*20+15]))
				case "int32_t", "int":
					chartData.Data = float64(variable.BytesToInt32(data[i*20+7 : i*20+15]))
				case "int64_t":
					chartData.Data = float64(variable.BytesToInt64(data[i*20+7 : i*20+15]))
				case "float":
					chartData.Data = float64(variable.BytesToFloat32(data[i*20+7 : i*20+15]))
				case "double":
					chartData.Data = float64(variable.BytesToFloat64(data[i*20+7 : i*20+15]))
				default:
					chartData.Data = 0
				}
				break
			}
		}

		if chartData.Name != "" {
			x.Variables = append(x.Variables, chartData)
		}
	}
}

func GetProjectVariables(x *variable.ListProjectT, data []byte) {
	reg := regexp.MustCompile(`(0x[0-9a-f]{8})\s{2}(0x[0-9a-f]+)\s+((\*\s)?[a-zA-Z0-9_\.]+)\s+([a-zA-Z0-9_\.\s]+?)[\n|\r]`)
	match := reg.FindAllStringSubmatch(string(data), -1)

	x.Variables = nil
	for _, v := range match {
		x.Variables = append(x.Variables, variable.ToProjectT{Addr: v[1], Size: v[2], Name: v[3], Type: v[5]})
	}
}
