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
		to[Write].Lock()
		defer to[Write].Lock()
		NewToModi := to[Write].m
		for k, v := range to[Write].m {
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
		to[Write].m = NewToModi
	}
}
