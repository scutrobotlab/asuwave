package variable

import (
	"encoding/json"
	"sync"

	"github.com/scutrobotlab/asuwave/pkg/jsonfile"
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

func SetAll(o Mod, v map[uint32]T) {
	to[o].Lock() // 锁保护
	defer to[o].Unlock()
	to[o].m = v
	jsonfile.Save(jsonPath[o], to[o].m)
}

// 以json格式获取所有Mod变量
func GetAll(o Mod) ([]byte, error) {
	to[o].RLock() // 锁保护
	defer to[o].RUnlock()
	return json.Marshal(to[o].m)
}

func Get(o Mod, k uint32) (T, bool) { //从map中读取一个值
	to[o].RLock()
	defer to[o].RUnlock()
	v, existed := to[o].m[k] // 在锁的保护下从map中读取
	return v, existed
}

func Set(o Mod, k uint32, v T) { // 设置一个键值对
	to[o].Lock() // 锁保护
	defer to[o].Unlock()
	to[o].m[k] = v
	jsonfile.Save(jsonPath[o], to[o].m)
}

func Delete(o Mod, k uint32) { //删除一个键
	to[o].Lock() // 锁保护
	defer to[o].Unlock()
	delete(to[o].m, k)
	jsonfile.Save(jsonPath[o], to[o].m)
}
