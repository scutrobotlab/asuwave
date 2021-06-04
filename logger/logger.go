package logger

import (
	"log"
	"os"
	"strconv"
	"time"
)

var Log *log.Logger

func init() {
	if _, err := os.Stat("log"); os.IsNotExist(err) {
		err := os.Mkdir("log", 0755)
		if err != nil {
			panic(err)
		}
	}

	logpath := "log/asuwave-" + strconv.FormatInt(time.Now().Unix(), 10) + ".log"
	file, err := os.Create(logpath)
	if err != nil {
		panic(err)
	}

	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.Println("LogFile : " + logpath)
}
