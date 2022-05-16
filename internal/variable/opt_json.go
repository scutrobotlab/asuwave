package variable

import (
	"path"

	"github.com/golang/glog"
	"github.com/scutrobotlab/asuwave/internal/helper"
	"github.com/scutrobotlab/asuwave/pkg/jsonfile"
)

var (
	jsonPath = map[Mod]string{
		RD: path.Join(helper.AppConfigDir(), "vToRead.json"),
		WR: path.Join(helper.AppConfigDir(), "vToWrite.json"),
	}
	optSaveVarList  bool
	optUpdateByProj bool
)

// Only call by option.
func SetOptSaveVarList(v bool) {
	if optSaveVarList == v {
		glog.V(1).Infof("SaveVarList has set to %t, skip\n", v)
		return
	}
	glog.V(1).Infof("Set SaveVarList to %t\n", v)
	jsonfile.Save(jsonPath[RD], to[RD].m)
	jsonfile.Save(jsonPath[WR], to[WR].m)
	optSaveVarList = v
}

// Only call by option.
func SetOptUpdateByProj(v bool) {
	if optUpdateByProj == v {
		glog.V(1).Infof("UpdateByProj has set to %t, skip\n", v)
		return
	}
	glog.V(1).Infof("Set UpdateByProj to %t\n", v)
	optUpdateByProj = v
}

// Only call by option.
func GetOptSaveVarList() bool {
	return optSaveVarList
}

// Only call by option.
func GetOptUpdateByProj() bool {
	return optUpdateByProj
}

func JsonLoadAll() {
	to[RD].Lock()
	defer to[RD].Unlock()
	to[WR].Lock()
	defer to[WR].Unlock()

	jsonfile.Load(jsonPath[RD], &to[RD].m)
	jsonfile.Load(jsonPath[WR], &to[WR].m)
	glog.Infoln(jsonPath[RD], "load success.")
	glog.Infoln(jsonPath[WR], "load success.")

	jsonfile.Save(jsonPath[RD], to[RD].m)
	jsonfile.Save(jsonPath[WR], to[WR].m)
}
