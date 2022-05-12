package jsonfile

import (
	"encoding/json"
	"os"

	"github.com/golang/glog"
)

func Save(filename string, v interface{}) {
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
