package jsonfile

import (
	"encoding/json"
	"os"
	"path"

	"github.com/golang/glog"
	"github.com/scutrobotlab/asuwave/internal/helper"
	"github.com/scutrobotlab/asuwave/internal/variable"
)

var (
	configPath   = path.Join(helper.AppConfigDir(), "config.json")
	variablePath = map[variable.Mod]string{
		variable.Read:  path.Join(helper.AppConfigDir(), "vToRead.json"),
		variable.Write: path.Join(helper.AppConfigDir(), "vToWrite.json"),
	}
	fileWatchPath = path.Join(helper.AppConfigDir(), "vFileWatch.json")
)

func Load(filename string, v interface{}) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		glog.Infoln(filename, " unfound.")
		return
	}
	glog.Infoln(filename, "Found")
	data, err := os.ReadFile(filename)
	if err != nil {
		glog.Errorln(err.Error())
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		glog.Errorln(err.Error())
	}
}

func UpdateVar(o variable.Mod) {
	if !option.options.SaveVarList {
		os.Remove(variablePath[o])
		return
	}
	data, err := variable.GetAll(o)
	if err != nil {
		glog.Errorln(err.Error())
		return
	}
	err = os.WriteFile(variablePath[o], data, 0644)
	if err != nil {
		glog.Errorln(err.Error())
	}
	glog.Infoln(variablePath[o], "save success.")
}

func save(filename string, v interface{}) {
	jsonTxt, err := json.Marshal(v)
	if err != nil {
		glog.Errorln(err.Error())
	}
	err = os.WriteFile(filename, jsonTxt, 0644)
	if err != nil {
		glog.Errorln(err.Error())
	}
	glog.Infoln(filename, "save success.")
}
