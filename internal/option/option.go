package option

import (
	"os"
	"path"

	"github.com/scutrobotlab/asuwave/internal/helper"
	"github.com/scutrobotlab/asuwave/internal/variable"
	"github.com/scutrobotlab/asuwave/pkg/file"
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

	var toRead = map[uint32]variable.T{}
	var toModi = map[uint32]variable.T{}
	JsonLoad(vToReadFileName, toRead)
	JsonLoad(vToModiFileName, toModi)
	variable.SetAll(variable.Read, toRead)
	variable.SetAll(variable.Modi, toModi)

	var watchList []string
	JsonLoad(vToProjFileName, &watchList)
	for _, w := range watchList {
		file.ChFileWatch <- w
	}
}

func Refresh() {
	if CheckCanSave(SaveVariableRead) {
		jsonSaveVar(variable.Read, vToReadFileName)
	} else {
		os.Remove(vToReadFileName)
	}
	if CheckCanSave(SaveVariableModi) {
		jsonSaveVar(variable.Modi, vToModiFileName)
	} else {
		os.Remove(vToModiFileName)
	}
	if CheckCanSave(SaveVariableProj) {
		JsonSave(vToProjFileName, file.GetWatchList())
	} else {
		os.Remove(vToProjFileName)
	}
}

func Save() {
	JsonSave(configFileName, Config)
}
