package option

import (
	"encoding/json"
	"os"

	"github.com/golang/glog"
	"github.com/scutrobotlab/asuwave/internal/variable"
)

func JsonLoad(filename string, v interface{}) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
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

func jsonSaveVar(m variable.Mod, filename string) {
	data, err := variable.GetAll(m)
	if err != nil {
		glog.Errorln(err.Error())
		return
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		glog.Errorln(err.Error())
	}
}

func JsonSave(filename string, v interface{}) {
	jsonTxt, err := json.Marshal(v)
	if err != nil {
		glog.Errorln(err.Error())
	}
	err = os.WriteFile(filename, jsonTxt, 0644)
	if err != nil {
		glog.Errorln(err.Error())
	}
}
