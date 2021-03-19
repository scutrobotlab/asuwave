package option

import (
	"os"
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

var Config ConfigT

const configFileName = "config.json"

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
}

func Save() {
	JsonSave(configFileName, Config)
}
