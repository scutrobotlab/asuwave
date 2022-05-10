package file

import (
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/golang/glog"
	"github.com/scutrobotlab/asuwave/internal/variable"
)

var watcher *fsnotify.Watcher
var watchList []string

var ChFileModi chan string = make(chan string, 10)
var ChFileError chan string = make(chan string, 10)
var ChFileWatch chan string = make(chan string, 10)

func GetWatchList() []string {
	glog.V(2).Infoln("Get: ", watchList)
	return watchList
}

func RemoveWathcer() error {
	l := watcher.WatchList()
	for _, p := range l {
		err := watcher.Remove(p)
		if err != nil {
			return err
		}
	}
	glog.V(2).Infoln("clear watcher")
	return nil
}

func FileWatch() {
	var err error
	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		glog.Errorln(err.Error())
		return
	}
	defer watcher.Close()
	watchdog := 0
	lastEventName := ""
	for {
		select {
		case file := <-ChFileWatch:
			for _, f := range watcher.WatchList() {
				watcher.Remove(f)
			}
			glog.Infoln("watch: ", file)
			watchList = []string{file}
			watcher.Add(file)
		case event, ok := <-watcher.Events:
			if !ok {
				glog.Warningln("Event not ok")
				return
			}
			glog.V(2).Infoln("file event:", event)
			if event.Op&fsnotify.Write == fsnotify.Write {
				lastEventName = event.Name
				watchdog = 0
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				glog.Warningln("Error not ok")
				return
			}
			lastEventName = ""
			ChFileError <- err.Error()
			glog.Errorln("error:", err)
		default:
			if lastEventName != "" && watchdog < 10 {
				watchdog++
			} else if lastEventName != "" && watchdog == 10 {
				glog.Infoln("file write done:", lastEventName)

				file, err := os.Open(lastEventName)
				if err != nil {
					glog.Errorln("file open:", err)
					return
				}

				f, err := Check(file)
				if err != nil {
					glog.Errorln("file check:", err)
					return
				}
				defer f.Close()

				err = ReadVariable(&variable.ToProj, f)
				if err != nil {
					glog.Errorln("file read:", err)
					return
				}
				variable.UpdateByProj()
				ChFileModi <- lastEventName
				watchdog++
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}
