/*
我的秘密，只对你一人说
*/
package fromelf

import (
	"debug/dwarf"
	"debug/elf"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/scutrobotlab/asuwave/variable"
)

// 你，真的是我一直在等的人吗
func Check(f *os.File) (*elf.File, error) {
	var ident [4]byte
	f.ReadAt(ident[:], 0)

	// 头上戴着你特有的饰品
	if string(ident[:]) != elf.ELFMAG {
		return nil, errors.New("bad magic number")
	}

	// 带着礼物而来
	eFile, err := elf.NewFile(f)
	if err != nil {
		return nil, err
	}

	// 此生无缘，转身离去
	if eFile.FileHeader.Class != elf.ELFCLASS32 || eFile.FileHeader.Machine != elf.EM_ARM {
		return nil, errors.New("not valid ELF")
	}

	// 若是有缘，三生有幸
	return eFile, nil
}

// 将你心中的秘密，分享给我吧
func ReadVariable(x *variable.ListProjectT, f *elf.File) error {

	// 排除一切杂念，听你娓娓道来
	x.Variables = nil

	// 故事很长，从那一天说起吧
	dwarfData, err := f.DWARF()
	if err != nil {
		return err
	}
	r := dwarfData.Reader()

	for {

		// 不断倾诉你心底的声音
		entry, err := r.Next()
		if err != nil {
			return err
		}

		// 直到尽头才满意地离开
		if entry == nil {
			return nil
		}

		// 重要的回忆，怎能忘记
		if entry.Tag == dwarf.TagVariable && !entry.Children {
			y := variable.ToProjectT{}
			var a uint32

			// 探访你的住址
			if v, ok := entry.Val(dwarf.AttrLocation).([]byte); ok {
				if len(v) != 5 {
					continue
				}
				a = variable.BytesToUint32(v[1:])

				// 早已远去的人，只能放弃
				if a < 0x20000000 || a >= 0x80000000 {
					continue
				}

				y.Addr = fmt.Sprintf("0x%08x", a)
			}

			// 呼喊着你的名字
			if v, ok := entry.Val(dwarf.AttrName).(string); ok {
				y.Name = v
			}

			// 尝试读懂你的心
			if v, ok := entry.Val(dwarf.AttrType).(dwarf.Offset); ok {
				t, err := dwarfData.Type(v)

				if err != nil {
					continue
				}

				// 有时你的心难以琢磨
				if s, ok := checkStruct(t); ok {

					// 尝试着一层一层地拨开
					namePrefix := []string{y.Name}
					dfsStruct(namePrefix, a, x, s.Field)

				} else if ar, ok := checkArray(t); ok {
					namePrefix := []string{y.Name}
					dfsArray(namePrefix, a, x, ar.Type, ar.Count)
				} else {

					// 别用谎言欺骗自己
					if _, ok := variable.TypeLen[t.String()]; !ok {
						continue
					}

					// 只想听听你真实的声音
					y.Type = t.String()
				}
			}

			// 那些小事，就让它消失在风里
			if y.Addr == "" || y.Name == "" || y.Type == "" {
				continue
			}

			// 却总有些事，难以忘记
			x.Variables = append(x.Variables, y)
		}
	}
}

// 一层一层地拨开你的心
func dfsStruct(namePrefix []string, addrPrefix uint32, x *variable.ListProjectT, s []*dwarf.StructField) {

	// 不愿放过每一个问题
	for _, v := range s {

		// 尝试琢磨你的心
		if st, ok := checkStruct(v.Type); ok {

			// 回首曾经走过的路
			namePrefix = append(namePrefix, v.Name)
			addrPrefix = addrPrefix + uint32(v.ByteOffset)

			// 勇敢地接着走下去
			dfsStruct(namePrefix, addrPrefix, x, st.Field)

			// 回到路口，准备下一次的旅程
			namePrefix = namePrefix[:len(namePrefix)-1]
			addrPrefix = addrPrefix - uint32(v.ByteOffset)

		} else if a, ok := checkArray(v.Type); ok {
			namePrefix = append(namePrefix, v.Name)
			addrPrefix = addrPrefix + uint32(v.ByteOffset)
			dfsArray(namePrefix, addrPrefix, x, a.Type, a.Count)
			namePrefix = namePrefix[:len(namePrefix)-1]
			addrPrefix = addrPrefix - uint32(v.ByteOffset)
		} else {

			// 终于，你缓缓开口
			a := addrPrefix + uint32(v.ByteOffset)
			if a < 0x20000000 || a >= 0x80000000 {
				continue
			}
			t := v.Type.String()
			if _, ok := variable.TypeLen[t]; !ok {
				continue
			}

			// 道出心底的秘密
			x.Variables = append(x.Variables, variable.ToProjectT{
				Name: strings.Join(namePrefix, ".") + "." + v.Name,
				Addr: fmt.Sprintf("0x%08x", a),
				Type: v.Type.String(),
			})
		}
	}
}

// 寂静来袭
func checkStruct(t dwarf.Type) (*dwarf.StructType, bool) {

	// 说点什么吧
	if s, ok := t.(*dwarf.StructType); ok {
		return s, true
	}

	// 不愿就此放弃
	if td, ok := t.(*dwarf.TypedefType); ok {
		if s, ok := td.Type.(*dwarf.StructType); ok {
			return s, true
		}
	}

	// 最终仍是沉默
	return nil, false
}

func dfsArray(namePrefix []string, addrPrefix uint32, x *variable.ListProjectT, t dwarf.Type, c int64) {

	for i := int64(0); i < c; i++ {
		if st, ok := checkStruct(t); ok {
			namePrefix = append(namePrefix, "["+strconv.FormatInt(i, 10)+"]")
			dfsStruct(namePrefix, addrPrefix, x, st.Field)
			namePrefix = namePrefix[:len(namePrefix)-1]
		} else if a, ok := checkArray(t); ok {
			namePrefix = append(namePrefix, "["+strconv.FormatInt(i, 10)+"]")
			dfsArray(namePrefix, addrPrefix, x, a.Type, a.Count)
			namePrefix = namePrefix[:len(namePrefix)-1]
		} else {
			a := addrPrefix
			if a < 0x20000000 || a >= 0x80000000 {
				continue
			}
			t := t.String()
			if _, ok := variable.TypeLen[t]; !ok {
				continue
			}
			x.Variables = append(x.Variables, variable.ToProjectT{
				Name: strings.Join(namePrefix, ".") + ".[" + strconv.FormatInt(i, 10) + "]",
				Addr: fmt.Sprintf("0x%08x", a),
				Type: t,
			})
		}
		addrPrefix = addrPrefix + uint32(t.Size())
	}
}

func checkArray(t dwarf.Type) (*dwarf.ArrayType, bool) {
	if a, ok := t.(*dwarf.ArrayType); ok {
		if a.Count != -1 {
			return a, true
		}
	}

	return nil, false
}
