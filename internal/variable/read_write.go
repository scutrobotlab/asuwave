package variable

import (
	"encoding/json"
	"sync"
)

type Mod int

const (
	Read Mod = iota
	Write
)

type T struct {
	Board      uint8
	Name       string
	Type       string
	Addr       uint32
	Data       float64
	Tick       uint32
	Inputcolor string
	SignalGain float64
	SignalBias float64
}

type RWMap struct { // 一个读写锁保护的线程安全的map
	sync.RWMutex // 读写锁保护下面的map字段
	m            map[uint32]T
}

var to []RWMap = []RWMap{{
	m: make(map[uint32]T, 0),
}, {
	m: make(map[uint32]T, 0),
}}

func SetAll(m Mod, m map[uint32]T) {
	to[m].Lock() // 锁保护
	defer to[m].Unlock()
	to[m].m = m
}

// 以json格式获取所有Opt变量
func GetAll(m Mod) ([]byte, error) {
	to[m].RLock() // 锁保护
	defer to[m].RUnlock()
	return json.Marshal(to[m].m)
}

func Get(m Mod, k uint32) (T, bool) { //从map中读取一个值
	to[m].RLock()
	defer to[m].RUnlock()
	v, existed := to[m].m[k] // 在锁的保护下从map中读取
	return v, existed
}

func Set(m Mod, k uint32, v T) { // 设置一个键值对
	to[m].Lock() // 锁保护
	defer to[m].Unlock()
	to[m].m[k] = v
}

func Delete(m Mod, k uint32) { //删除一个键
	to[m].Lock() // 锁保护
	defer to[m].Unlock()
	delete(to[m].m, k)
}
