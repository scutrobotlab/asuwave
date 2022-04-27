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

	a := (n - r) % 20
	if a == 0 {
		a = 20
	}

	// 也许收获满满
	r = n - a
	return l, r + 1
}

func Convert_T_ToChartT(x variable.T) variable.ToChartT {
	return variable.ToChartT{
		Board: x.Board,
		Name:  x.Name,
		Data:  x.SignalGain*x.Data + x.SignalBias,
		Tick:  x.Tick,
	}
}

// 从茫茫 data 中，寻找我所挂念的列表 y ，记录在列表 x 中。
// 所有的 add 我都难以忘记，所有的 del 我都不愿提起
func MakeChartPack(y *variable.ListT, data []byte) (x variable.ListChartT, add variable.ListT, del variable.ListT) {
	x = variable.ListChartT{}
	add = variable.ListT{} // 有些变量，我难以忘记
	del = variable.ListT{} // 有些变量，我不愿提起
	dataList := variable.ListT{}
	for i := 0; i < len(data)/20; i++ {
		// 解开关于它的一切
		board := data[i*20]
		typelen := int(data[i*20+2])
		addr := variable.BytesToUint32(data[i*20+3 : i*20+7])
		tick := variable.BytesToUint32(data[i*20+15 : i*20+19])

		// 它是我要找的那个变量吗？
		if v, ok := (*y)[addr]; ok { // 是的，我还挂念着它
			switch v.Type {
			case "uint8_t":
				v.Data = float64(variable.BytesToUint8(data[i*20+7 : i*20+15]))
			case "uint16_t":
				v.Data = float64(variable.BytesToUint16(data[i*20+7 : i*20+15]))
			case "uint32_t":
				v.Data = float64(variable.BytesToUint32(data[i*20+7 : i*20+15]))
			case "uint64_t":
				v.Data = float64(variable.BytesToUint64(data[i*20+7 : i*20+15]))
			case "int8_t":
				v.Data = float64(variable.BytesToInt8(data[i*20+7 : i*20+15]))
			case "int16_t":
				v.Data = float64(variable.BytesToInt16(data[i*20+7 : i*20+15]))
			case "int32_t", "int":
				v.Data = float64(variable.BytesToInt32(data[i*20+7 : i*20+15]))
			case "int64_t":
				v.Data = float64(variable.BytesToInt64(data[i*20+7 : i*20+15]))
			case "float":
				v.Data = float64(variable.BytesToFloat32(data[i*20+7 : i*20+15]))
			case "double":
				v.Data = float64(variable.BytesToFloat64(data[i*20+7 : i*20+15]))
			default:
				v.Data = 0
			}
			v.Tick = tick // 同步它的心跳
			dataList[addr] = v

			x = append(x, Convert_T_ToChartT(v))
		} else { // 不是的，请忘了它
			del[addr] = variable.T{
				Board: board,
				Type:  variable.LenType[typelen], // 垃圾代码：由于variable.T需要Type字段（删除变量需要TypeLen），但是mcu只会反馈TypeLen，所以强行安排了一个Type。
				Addr:  addr,
				Tick:  tick,
			}
		}
	}

	// 我所挂念的，它们都还在吗
	for _, v := range *y {
		if _, ok := dataList[v.Addr]; !ok {
			// 我很想它，下次请别忘记
			add[v.Addr] = v
		}
	}
	return
}
