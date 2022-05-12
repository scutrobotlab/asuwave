package variable

import (
	"path"

	"github.com/golang/glog"
	"github.com/scutrobotlab/asuwave/internal/helper"
	"github.com/scutrobotlab/asuwave/pkg/jsonfile"
)

var (
	jsonPath = map[Mod]string{
		Read:  path.Join(helper.AppConfigDir(), "vToRead.json"),
		Write: path.Join(helper.AppConfigDir(), "vToWrite.json"),
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
	jsonfile.Save(jsonPath[Read], to[Read].m)
	jsonfile.Save(jsonPath[Write], to[Write].m)
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
	to[Read].Lock()
	defer to[Read].Unlock()
	to[Write].Lock()
	defer to[Write].Unlock()

	jsonfile.Load(jsonPath[Read], &to[Read].m)
	jsonfile.Load(jsonPath[Write], &to[Write].m)
	glog.Infoln(jsonPath[Read], "load success.")
	glog.Infoln(jsonPath[Write], "load success.")

	jsonfile.Save(jsonPath[Read], to[Read].m)
	jsonfile.Save(jsonPath[Write], to[Write].m)
}
