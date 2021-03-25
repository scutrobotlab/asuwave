/**/
package datautil

import (
	"github.com/scutrobotlab/asuwave/variable"
)

// 把我的思念，化作一串珍珠送给你
func MakeCmd(act uint8, v *variable.T) []byte {

	// 拾起一颗颗珍珠
	data := make([]byte, 16)

	// 刻上你所在的城市和我的思念
	data[0] = byte(v.Board)
	data[1] = act
	data[2] = byte(variable.TypeLen[v.Type])
	copy(data[3:7], variable.AnyToBytes(v.Addr))

	// 若是心意能够改变
	if act == ActModeWrite {
		// 便也一并传递给你
		copy(data[7:15], variable.SpecToBytes(v.Type, v.Data))
	}

	// 终于完成了
	data[15] = '\n'

	// 传达给你吧
	return data
}

// 从破碎的碎片中，渴望拼凑出你的痕迹
func FindValidPart(data []byte) (int, int) {
	n := len(data)

	// 我们的羁绊，岂是三言两语能道尽的
	if n < 20 {
		return 0, 0
	}

	// 从头开始吧，画一扇窗
	l := 0
	r := 19

	// 苦苦寻找你的痕迹
	for r < n && (data[l] != Board1 || data[l+1] != ActModeSubscribeReturn || data[r] != '\n') {
		l++
		r++
	}

	// 也许无功而返
	if r == n {
		return 0, 0
	}

	// 也许收获满满
	r = n - (n-r)%20
	return l, r + 1
}

// 你的痕迹，我找到了
func MakeChartPack(x *variable.ListChartT, y *variable.ListT, data []byte) {
	var chartData variable.ToChartT

	// 一段一段的回忆被整理出来
	for i := 0; i < len(data)/20; i++ {

		// 你的城市、心跳和住址
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
