//go:build release

package logger

import (
	"log"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/scutrobotlab/asuwave/internal/helper"
)

var Log *log.Logger

func init() {
	logdir := path.Join(helper.AppConfigDir(), "log")
	if _, err := os.Stat(logdir); os.IsNotExist(err) {
		err := os.Mkdir(logdir, 0755)
		if err != nil {
			panic(err)
		}
	}

	logpath := path.Join(logdir, "asuwave-"+strconv.FormatInt(time.Now().Unix(), 10)+".log")
	file, err := os.Create(logpath)
	if err != nil {
		panic(err)
	}

	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.Println(helper.GetVersion())
	Log.Println("LogFile : " + logpath)
}
