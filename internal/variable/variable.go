package variable

import (
	"path"
	"strconv"

	"github.com/golang/glog"
	"github.com/scutrobotlab/asuwave/internal/helper"
)

var TypeLen = map[string]int{
	"uint8_t":  1,
	"uint16_t": 2,
	"uint32_t": 4,
	"uint64_t": 8,
	"int8_t":   1,
	"int16_t":  2,
	"int32_t":  4,
	"int64_t":  8,
	"int":      4,
	"float":    4,
	"double":   8,
}

var (
	vToReadFileName = path.Join(helper.AppConfigDir(), "vToRead.json")
	vToModiFileName = path.Join(helper.AppConfigDir(), "vToModi.json")
	vToProjFileName = path.Join(helper.AppConfigDir(), "vToProj.json")
)

// 通过Proj的变量名更新Read和Modi的地址和类型
func UpdateByProj() {
	toProj.RLock()
	defer toProj.RUnlock()
	{
		to[Read].Lock()
		defer to[Read].Lock()
		NewToRead := to[Read].m
		for k, v := range to[Read].m {
			if p, ok := toProj.m[v.Name]; ok {
				addr, err := strconv.ParseUint(p.Addr, 16, 32)
				if err != nil {
					glog.Errorln(err.Error())
					continue
				}
				v.Addr = uint32(addr)
				v.Type = p.Type
				delete(NewToRead, k)
				NewToRead[v.Addr] = v
			}
		}
		to[Read].m = NewToRead
	}
	{
		to[Modi].Lock()
		defer to[Modi].Lock()
		NewToModi := to[Modi].m
		for k, v := range to[Modi].m {
			if p, ok := toProj.m[v.Name]; ok {
				addr, err := strconv.ParseUint(p.Addr, 16, 32)
				if err != nil {
					glog.Errorln(err.Error())
					continue
				}
				v.Addr = uint32(addr)
				v.Type = p.Type
				delete(NewToModi, k)
				NewToModi[v.Addr] = v
			}
		}
		to[Modi].m = NewToModi
	}
}
