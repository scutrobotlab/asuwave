package fromelf

import (
	"debug/dwarf"
	"debug/elf"
	"errors"
	"fmt"
	"os"

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
			if v, ok := entry.Val(dwarf.AttrLocation).([]byte); ok {
				if len(v) != 5 {
					continue
				}
				a := variable.BytesToUint32(v[1:])
				if a < 0x20000000 || a >= 0x80000000 {
					continue
				}
				y.Addr = fmt.Sprintf("0x%08x", a)
			}
			if v, ok := entry.Val(dwarf.AttrName).(string); ok {
				y.Name = v
			}
			if v, ok := entry.Val(dwarf.AttrType).(dwarf.Offset); ok {
				if t, err := dwarfData.Type(v); err == nil {
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
