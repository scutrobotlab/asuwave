package variable

import (
	"log"
	"path"
	"strconv"

	"github.com/scutrobotlab/asuwave/helper"
)

var LenType = map[int]string{
	1: "uint8_t",
	2: "uint16_t",
	4: "uint32_t",
	8: "uint64_t",
}

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

type ListT map[uint32]T

var ToRead ListT = ListT{}
var ToModi ListT = ListT{}

type ToProjectT struct {
	Addr string
	Name string
	Type string
}
type ListProjectT map[string]ToProjectT

var ToProj ListProjectT = ListProjectT{}

type ToChartT struct {
	Board uint8
	Name  string
	Data  float64
	Tick  uint32
}

type ListChartT []ToChartT

var (
	vToReadFileName = path.Join(helper.AppConfigDir(), "vToRead.json")
	vToModiFileName = path.Join(helper.AppConfigDir(), "vToModi.json")
	vToProjFileName = path.Join(helper.AppConfigDir(), "vToProj.json")
)

func Update() {
	{
		NewToRead := ToRead
		for k, v := range ToRead {
			if p, ok := ToProj[v.Name]; ok {
				addr, err := strconv.ParseUint(p.Addr, 16, 32)
				if err != nil {
					log.Println(err.Error())
					continue
				}
				v.Addr = uint32(addr)
				v.Type = p.Type
				delete(NewToRead, k)
				NewToRead[v.Addr] = v
			}
		}
		ToRead = NewToRead
	}
	{
		NewToModi := ToModi
		for k, v := range ToModi {
			if p, ok := ToProj[v.Name]; ok {
				addr, err := strconv.ParseUint(p.Addr, 16, 32)
				if err != nil {
					log.Println(err.Error())
					continue
				}
				v.Addr = uint32(addr)
				v.Type = p.Type
				delete(NewToModi, k)
				NewToModi[v.Addr] = v
			}
		}
		ToModi = NewToModi
	}
}
