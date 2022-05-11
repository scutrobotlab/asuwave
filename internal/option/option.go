package option

import (
	"os"
	"path"

	"github.com/scutrobotlab/asuwave/internal/helper"
	"github.com/scutrobotlab/asuwave/internal/variable"
	"github.com/scutrobotlab/asuwave/pkg/elffile"
)

type OptType struct {
	LogLevel     int
	SaveVarList  bool
	SaveFilePath bool
	UpdateByProj bool
}

var options OptType

var (
	configJson    = path.Join(helper.AppConfigDir(), "config.json")
	vToReadJson   = path.Join(helper.AppConfigDir(), "vToRead.json")
	vToWriteJson  = path.Join(helper.AppConfigDir(), "vToWrite.json")
	fileWatchJson = path.Join(helper.AppConfigDir(), "vFileWatch.json")
)

func Load() {
	if _, err := os.Stat(configJson); !os.IsNotExist(err) {
		JsonLoad(configJson, options)
	}

	var toRead = map[uint32]variable.T{}
	var toWrite = map[uint32]variable.T{}
	JsonLoad(vToReadJson, toRead)
	JsonLoad(vToWriteJson, toWrite)
	variable.SetAll(variable.Read, toRead)
	variable.SetAll(variable.Write, toWrite)

	var watchList []string
	JsonLoad(fileWatchJson, watchList)
	for _, w := range watchList {
		elffile.ChFileWatch <- w
	}

	jsonSaveVar(variable.Read, vToReadJson)
	jsonSaveVar(variable.Write, vToWriteJson)
	JsonSave(fileWatchJson, elffile.GetWatchList())
	JsonSave(configJson, options)
}

func GetAll() OptType {
	return options
}

func SetLogLevel(v int) {
	if options.LogLevel == v {
		return
	}
	options.LogLevel = v
	JsonSave(configJson, options)
}

func SetSaveVarList(v bool) {
	if options.SaveVarList == v {
		return
	}
	if v {
		jsonSaveVar(variable.Read, vToReadJson)
		jsonSaveVar(variable.Write, vToWriteJson)
	} else {
		os.Remove(vToReadJson)
		os.Remove(vToWriteJson)
	}
	options.SaveVarList = v
	JsonSave(configJson, options)
}

func SetSaveFilePath(v bool) {
	if options.SaveFilePath == v {
		return
	}
	if options.SaveFilePath {
		JsonSave(fileWatchJson, elffile.GetWatchList())
	} else {
		os.Remove(fileWatchJson)
	}
	options.SaveFilePath = v
	JsonSave(configJson, options)
}

func SetUpdateByProj(v bool) {
	if options.UpdateByProj == v {
		return
	}
	options.UpdateByProj = v
	JsonSave(configJson, options)
}
