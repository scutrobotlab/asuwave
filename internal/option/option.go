package option

import (
	"os"
	"path"

	"github.com/scutrobotlab/asuwave/internal/helper"
	"github.com/scutrobotlab/asuwave/internal/variable"
	"github.com/scutrobotlab/asuwave/pkg/elffile"
)

var options struct {
	LogLevel     int
	SaveVarList  bool
	SaveFilePath bool
	UpdateByProj bool
}

var (
	configJson     = path.Join(helper.AppConfigDir(), "config.json")
	vToReadJson    = path.Join(helper.AppConfigDir(), "vToRead.json")
	vToWriteJson   = path.Join(helper.AppConfigDir(), "vToWrite.json")
	vFileWatchJson = path.Join(helper.AppConfigDir(), "vFileWatch.json")
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
	JsonLoad(vFileWatchJson, watchList)
	for _, w := range watchList {
		elffile.ChFileWatch <- w
	}
}

func SetLogLevel(v int) {
	if options.LogLevel == v {
		return
	}
	options.LogLevel = v
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
}

func SetSaveFilePath(v bool) {
	if options.SaveFilePath == v {
		return
	}
	if options.SaveFilePath {
		JsonSave(vFileWatchJson, elffile.GetWatchList())
	} else {
		os.Remove(vFileWatchJson)
	}
	options.SaveFilePath = v
}

func SetUpdateByProj(v bool) {
	if options.UpdateByProj == v {
		return
	}
	options.UpdateByProj = v
}

func Save() {
	JsonSave(configJson, options)
}
