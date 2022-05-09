package variable

import (
	"encoding/json"
	"sync"
)

type ProjT struct {
	Addr string
	Name string
	Type string
}

type ProjMapType struct { // 一个读写锁保护的线程安全的map
	sync.RWMutex // 读写锁保护下面的map字段
	m            map[string]ProjT
}

var toProj ProjMapType = ProjMapType{
	m: make(map[string]ProjT, 0),
}

func SetAllProj(m map[string]ProjT) {
	toProj.Lock() // 锁保护
	defer toProj.Unlock()
	toProj.m = m
}

// 以json格式获取所有Proj变量
func GetAllProj() ([]byte, error) {
	toProj.RLock() // 锁保护
	defer toProj.RUnlock()
	return json.Marshal(toProj.m)
}

func GetProj(k string) (ProjT, bool) { //从map中读取一个值
	toProj.RLock()
	defer toProj.RUnlock()
	v, existed := toProj.m[k] // 在锁的保护下从map中读取
	return v, existed
}

func SetProj(k string, v ProjT) { // 设置一个键值对
	toProj.Lock() // 锁保护
	defer toProj.Unlock()
	toProj.m[k] = v
}

func DeleteProj(k string) { //删除一个键
	toProj.Lock() // 锁保护
	defer toProj.Unlock()
	delete(toProj.m, k)
}
