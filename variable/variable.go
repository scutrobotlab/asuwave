package variable

import (
	"os"

	"github.com/scutrobotlab/asuwave/option"
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

type T struct {
	Board uint8
	Name  string
	Type  string
	Addr  uint32
	Data  float64
	Tick  uint32
}
type ListT struct {
	Variables []T
}

var ToRead ListT
var ToModi ListT

type ToProjectT struct {
	Addr string
	Size string
	Name string
	Type string
}
type ListProjectT struct {
	Variables []ToProjectT
}

var ToProj ListProjectT

type ToChartT struct {
	Board uint8
	Name  string
	Data  float64
	Tick  uint32
}

type ListChartT struct {
	Variables []ToChartT
}

const vToReadFileName = "vToRead.json"
const vToModiFileName = "vToModi.json"
const vToProjFileName = "vToProj.json"

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
