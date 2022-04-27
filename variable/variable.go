package variable

import (
	"os"
	"path"

	"github.com/scutrobotlab/asuwave/helper"
	"github.com/scutrobotlab/asuwave/option"
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

func Load() {
	option.JsonLoad(vToReadFileName, &ToRead)
	option.JsonLoad(vToModiFileName, &ToModi)
	option.JsonLoad(vToProjFileName, &ToProj)
}

func Refresh() {
	if option.CheckCanSave(option.SaveVariableRead) {
		option.JsonSave(vToReadFileName, &ToRead)
	} else {
		os.Remove(vToReadFileName)
	}
	if option.CheckCanSave(option.SaveVariableModi) {
		option.JsonSave(vToModiFileName, &ToModi)
	} else {
		os.Remove(vToModiFileName)
	}
	if option.CheckCanSave(option.SaveVariableProj) {
		option.JsonSave(vToProjFileName, &ToProj)
	} else {
		os.Remove(vToProjFileName)
	}
}
