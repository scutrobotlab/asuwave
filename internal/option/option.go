package option

import (
	"flag"
	"os"
	"path"
	"strconv"

	"github.com/golang/glog"
	"github.com/scutrobotlab/asuwave/internal/helper"
	"github.com/scutrobotlab/asuwave/internal/variable"
	"github.com/scutrobotlab/asuwave/pkg/elffile"
	"github.com/scutrobotlab/asuwave/pkg/jsonfile"
)

var (
	logLevel     int
	saveFilePath bool
)

var (
	optionPath    = path.Join(helper.AppConfigDir(), "option.json")
	fileWatchPath = path.Join(helper.AppConfigDir(), "FileWatch.json")
)

type OptT struct {
	LogLevel     int
	SaveFilePath bool
	SaveVarList  bool
	UpdateByProj bool
}

func Get() OptT {
	return OptT{
		LogLevel:     logLevel,
		SaveFilePath: saveFilePath,
		SaveVarList:  variable.GetOptSaveVarList(),
		UpdateByProj: variable.GetOptUpdateByProj(),
	}
}

func Load() {
	var opt OptT
	jsonfile.Load(optionPath, &opt)
	variable.SetOptSaveVarList(opt.SaveVarList)
	variable.SetOptUpdateByProj(opt.UpdateByProj)

	var watchList []string
	jsonfile.Load(fileWatchPath, &watchList)
	for _, w := range watchList {
		elffile.ChFileWatch <- w
	}

	jsonfile.Save(fileWatchPath, elffile.GetWatchList())
	jsonfile.Save(optionPath, opt)
}

func SetLogLevel(v int) {
	if logLevel == v {
		glog.V(1).Infof("LogLevel has set to %d, skip\n", v)
		return
	}
	glog.V(1).Infof("Set LogLevel to %d\n", v)
	logLevel = v
	if err := flag.Set("v", strconv.Itoa(v)); err != nil {
		glog.Errorln(err.Error())
	}
	jsonfile.Save(optionPath, Get())
}

func SetSaveFilePath(v bool) {
	if saveFilePath == v {
		glog.V(1).Infof("SaveFilePath has set to %t, skip\n", v)
		return
	}
	glog.V(1).Infof("Set SaveFilePath to %t\n", v)
	if v {
		jsonfile.Save(fileWatchPath, elffile.GetWatchList())
	} else {
		os.Remove(fileWatchPath)
	}
	saveFilePath = v
	jsonfile.Save(optionPath, Get())
}

func SetSaveVarList(v bool) {
	variable.SetOptSaveVarList(v)
	jsonfile.Save(optionPath, Get())
}

func SetUpdateByProj(v bool) {
	variable.SetOptUpdateByProj(v)
	jsonfile.Save(optionPath, Get())
}
