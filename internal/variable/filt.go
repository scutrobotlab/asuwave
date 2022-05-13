package variable

type CmdT struct {
	Board   uint8
	TypeLen int
	Addr    uint32
}

// 从茫茫 data 中，寻找我所挂念的 to[Read] ，记录在列表 chart 中。
// 所有的 add 我都难以忘记，所有的 del 我都不愿提起
func Filt(data []byte) (chart []ChartT, add []CmdT, del []CmdT) {
	to[Read].RLock()
	defer to[Read].RUnlock()

	chart = []ChartT{}
	add = []CmdT{} // 有些变量，我难以忘记
	del = []CmdT{} // 有些变量，我不愿提起

	addrs := map[uint32]bool{}

	for i := 0; i < len(data)/20; i++ {
		// 解开关于它的一切
		board := data[i*20]
		typelen := int(data[i*20+2])
		addr := BytesToUint32(data[i*20+3 : i*20+7])
		tick := BytesToUint32(data[i*20+15 : i*20+19])

		addrs[addr] = true

		// 它是我要找的那个变量吗？
		if v, ok := to[Read].m[addr]; ok { // 是的，我还挂念着它
			switch v.Type {
			case "uint8_t":
				v.Data = float64(BytesToUint8(data[i*20+7 : i*20+15]))
			case "uint16_t":
				v.Data = float64(BytesToUint16(data[i*20+7 : i*20+15]))
			case "uint32_t":
				v.Data = float64(BytesToUint32(data[i*20+7 : i*20+15]))
			case "uint64_t":
				v.Data = float64(BytesToUint64(data[i*20+7 : i*20+15]))
			case "int8_t":
				v.Data = float64(BytesToInt8(data[i*20+7 : i*20+15]))
			case "int16_t":
				v.Data = float64(BytesToInt16(data[i*20+7 : i*20+15]))
			case "int32_t", "int":
				v.Data = float64(BytesToInt32(data[i*20+7 : i*20+15]))
			case "int64_t":
				v.Data = float64(BytesToInt64(data[i*20+7 : i*20+15]))
			case "float":
				v.Data = float64(BytesToFloat32(data[i*20+7 : i*20+15]))
			case "double":
				v.Data = float64(BytesToFloat64(data[i*20+7 : i*20+15]))
			default:
				v.Data = 0
			}
			v.Tick = tick // 同步它的心跳
			chart = append(chart, ChartT{
				Board: v.Board,
				Name:  v.Name,
				Data:  v.SignalGain*v.Data + v.SignalBias,
				Tick:  v.Tick,
			})
		} else { // 不是的，请忘了它
			del = append(del, CmdT{
				Board:   board,
				TypeLen: typelen,
				Addr:    addr,
			})
		}
	}

	// 我所挂念的，它们都还在吗
	for _, v := range to[Read].m {
		if _, ok := addrs[v.Addr]; !ok {
			// 我很想它，下次请别忘记
			add = append(add, CmdT{
				Board:   v.Board,
				TypeLen: TypeLen[v.Type],
				Addr:    v.Addr,
			})
		}
	}
	return
}
