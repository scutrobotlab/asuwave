package option

import (
	"os"
	"path"

	"github.com/scutrobotlab/asuwave/fromelf"
	"github.com/scutrobotlab/asuwave/helper"
	"github.com/scutrobotlab/asuwave/variable"
)

const (
	SaveVariableProj = 1
	SaveVariableRead = 2
	SaveVariableModi = 4
)

type ConfigT struct {
	Save int
	Port int
}

var (
	vToReadFileName = path.Join(helper.AppConfigDir(), "vToRead.json")
	vToModiFileName = path.Join(helper.AppConfigDir(), "vToModi.json")
	vToProjFileName = path.Join(helper.AppConfigDir(), "vToProj.json")
)

var Config ConfigT

var configFileName = path.Join(helper.AppConfigDir(), "config.json")

func CheckCanSave(s int) bool {
	return s&Config.Save == s
}

func Load() {
	if _, err := os.Stat(configFileName); os.IsNotExist(err) {
		Config.Save = 7
		Config.Port = 8000
	} else {
		JsonLoad(configFileName, &Config)
	}

	JsonLoad(vToReadFileName, &variable.ToRead)
	JsonLoad(vToModiFileName, &variable.ToModi)
	var watchList []string
	JsonLoad(vToProjFileName, &watchList)
	for _, w := range watchList {
		fromelf.ChFileWatch <- w
	}
}

func Refresh() {
	if CheckCanSave(SaveVariableRead) {
		JsonSave(vToReadFileName, &variable.ToRead)
	} else {
		os.Remove(vToReadFileName)
	}
	if CheckCanSave(SaveVariableModi) {
		JsonSave(vToModiFileName, &variable.ToModi)
	} else {
		os.Remove(vToModiFileName)
	}
	if CheckCanSave(SaveVariableProj) {
		JsonSave(vToProjFileName, fromelf.GetWatchList())
	} else {
		os.Remove(vToProjFileName)
	}
}

func Save() {
	JsonSave(configFileName, Config)
}
