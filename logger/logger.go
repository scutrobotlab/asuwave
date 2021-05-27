package logger

import (
	"log"
	"os"
	"strconv"
	"time"
)

var Log *log.Logger

func init() {
	logpath := "asuwave-" + strconv.FormatInt(time.Now().Unix(), 10) + ".log"
	file, err := os.Create(logpath)
	if err != nil {
		panic(err)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.Println("LogFile : " + logpath)
}
