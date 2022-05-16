package variable

import (
	"strconv"

	"github.com/golang/glog"
)

// 通过Proj的变量名更新Read和Write的地址和类型
func UpdateByProj() {
	toProj.RLock()
	defer toProj.RUnlock()
	{
		to[RD].Lock()
		defer to[RD].Lock()
		NewToRead := to[RD].m
		for k, v := range to[RD].m {
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
		to[RD].m = NewToRead
	}
	{
		to[WR].Lock()
		defer to[WR].Lock()
		NewToModi := to[WR].m
		for k, v := range to[WR].m {
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
		to[WR].m = NewToModi
	}
}
