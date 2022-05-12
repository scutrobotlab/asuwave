package option

import (
	"flag"
	"os"
	"strconv"

	"github.com/golang/glog"
	"github.com/scutrobotlab/asuwave/internal/jsonfile"
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

func Load() {
	jsonfile.Load(configJson, &options)

	var toRead = map[uint32]variable.T{}
	var toWrite = map[uint32]variable.T{}
	jsonfile.Load(variableJson[variable.Read], &toRead)
	jsonfile.Load(variableJson[variable.Write], &toWrite)
	variable.SetAll(variable.Read, toRead)
	variable.SetAll(variable.Write, toWrite)

	var watchList []string
	jsonfile.Load(fileWatchJson, &watchList)
	for _, w := range watchList {
		elffile.ChFileWatch <- w
	}

	jsonfile.UpdateVar(variable.Read)
	jsonfile.UpdateVar(variable.Write)

	jsonfile.Save(fileWatchJson, elffile.GetWatchList())
	jsonfile.Save(configJson, options)
}

func GetAll() OptType {
	return options
}

func SetLogLevel(v int) {
	if options.LogLevel == v {
		glog.V(1).Infof("LogLevel has set to %d, skip\n", v)
		return
	}
	glog.V(1).Infof("Set LogLevel to %t\n", v)
	options.LogLevel = v
	if err := flag.Set("v", strconv.Itoa(v)); err != nil {
		glog.Errorln(err.Error())
	}
	jsonfile.Save(configJson, options)
}

func SetSaveVarList(v bool) {
	if options.SaveVarList == v {
		glog.V(1).Infof("SaveVarList has set to %t, skip\n", v)
		return
	}
	glog.V(1).Infof("Set SaveVarList to %t\n", v)
	jsonfile.UpdateVar(variable.Read)
	jsonfile.UpdateVar(variable.Write)
	options.SaveVarList = v
	jsonfile.Save(configJson, options)
}

func SetSaveFilePath(v bool) {
	if options.SaveFilePath == v {
		glog.V(1).Infof("SaveFilePath has set to %t, skip\n", v)
		return
	}
	glog.V(1).Infof("Set SaveFilePath to %t\n", v)
	if options.SaveFilePath {
		jsonfile.Save(fileWatchJson, elffile.GetWatchList())
	} else {
		os.Remove(fileWatchJson)
	}
	options.SaveFilePath = v
	jsonfile.Save(configJson, options)
}

func SetUpdateByProj(v bool) {
	if options.UpdateByProj == v {
		glog.V(1).Infof("UpdateByProj has set to %t, skip\n", v)
		return
	}
	glog.V(1).Infof("Set UpdateByProj to %t\n", v)
	options.UpdateByProj = v
	jsonfile.Save(configJson, options)
}
