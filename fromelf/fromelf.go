package fromelf

import (
	"debug/dwarf"
	"debug/elf"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/scutrobotlab/asuwave/variable"
)

func Check(f *os.File) (*elf.File, error) {
	var ident [4]byte
	f.ReadAt(ident[:], 0)
	if string(ident[:]) != elf.ELFMAG {
		return nil, errors.New("bad magic number")
	}
	eFile, err := elf.NewFile(f)
	if err != nil {
		return nil, err
	}
	if eFile.FileHeader.Class != elf.ELFCLASS32 || eFile.FileHeader.Machine != elf.EM_ARM {
		return nil, errors.New("not valid ELF")
	}
	return eFile, nil
}

func ReadVariable(x *variable.ListProjectT, f *elf.File) error {
	var entry *dwarf.Entry
	var dwarfData *dwarf.Data
	var err error
	var y variable.ToProjectT
	x.Variables = nil

	dwarfData, err = f.DWARF()
	if err != nil {
		return err
	}
	r := dwarfData.Reader()

	for {
		y = variable.ToProjectT{}
		entry, err = r.Next()
		if err != nil {
			return err
		}
		if entry == nil {
			return nil
		}
		if entry.Tag == dwarf.TagVariable && !entry.Children {
			var a uint32
			if v, ok := entry.Val(dwarf.AttrLocation).([]byte); ok {
				if len(v) != 5 {
					continue
				}
				a = variable.BytesToUint32(v[1:])
				if a < 0x20000000 || a >= 0x80000000 {
					continue
				}
				y.Addr = fmt.Sprintf("0x%08x", a)
			}

			if v, ok := entry.Val(dwarf.AttrName).(string); ok {
				y.Name = v
			}
			if v, ok := entry.Val(dwarf.AttrType).(dwarf.Offset); ok {
				t, err := dwarfData.Type(v)

				if err != nil {
					continue
				}

				if s, ok := checkStruct(t); ok {
					var namePrefix []string
					namePrefix = append(namePrefix, y.Name)
					dfs(namePrefix, a, x, s.Field)
				} else {
					y.Type = t.String()
				}
			}

			if y.Addr == "" || y.Name == "" || y.Type == "" {
				continue
			}

			x.Variables = append(x.Variables, y)
		}
	}
}

func dfs(namePrefix []string, addrPrefix uint32, x *variable.ListProjectT, s []*dwarf.StructField) {
	for _, v := range s {
		if st, ok := checkStruct(v.Type); ok {
			namePrefix = append(namePrefix, v.Name)
			addrPrefix = addrPrefix + uint32(v.ByteOffset)
			dfs(namePrefix, addrPrefix, x, st.Field)
			namePrefix = namePrefix[:len(namePrefix)-1]
			addrPrefix = addrPrefix - uint32(v.ByteOffset)
		} else {
			a := addrPrefix + uint32(v.ByteOffset)
			if a < 0x20000000 || a >= 0x80000000 {
				continue
			}
			x.Variables = append(x.Variables, variable.ToProjectT{
				Name: strings.Join(namePrefix, ".") + "." + v.Name,
				Addr: fmt.Sprintf("0x%08x", a),
				Type: v.Type.String(),
			})
		}
	}
}

func checkStruct(t dwarf.Type) (*dwarf.StructType, bool) {
	if td, ok := t.(*dwarf.TypedefType); ok {
		if s, ok := td.Type.(*dwarf.StructType); ok {
			return s, true
		}
	}
	return nil, false
}
