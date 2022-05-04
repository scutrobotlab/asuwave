package fromelf

import (
	"log"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/scutrobotlab/asuwave/variable"
)

var watcher *fsnotify.Watcher

var ChFileModi chan string = make(chan string, 10)
var ChFileError chan string = make(chan string, 10)
var ChFileWatch chan string = make(chan string, 10)

func GetWatchList() []string {
	return watcher.WatchList()
}

func RemoveWathcer() error {
	l := watcher.WatchList()
	for _, p := range l {
		err := watcher.Remove(p)
		if err != nil {
			return err
		}
	}
	return nil
}

func FileWatch() {
	var err error
	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
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
			watcher.Add(file)
		case event, ok := <-watcher.Events:
			if !ok {
				log.Println("Event not ok")
				return
			}
			log.Println("file event:", event)
			if event.Op&fsnotify.Write == fsnotify.Write {
				lastEventName = event.Name
				watchdog = 0
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				log.Println("Error not ok")
				return
			}
			lastEventName = ""
			ChFileError <- err.Error()
			log.Println("error:", err)
		default:
			if lastEventName != "" && watchdog < 10 {
				watchdog++
			} else if lastEventName != "" && watchdog == 10 {
				log.Println("file write done:", lastEventName)

				file, err := os.Open(lastEventName)
				if err != nil {
					log.Println("file open:", err)
					return
				}

				f, err := Check(file)
				if err != nil {
					log.Println("file check:", err)
					return
				}
				defer f.Close()

				err = ReadVariable(&variable.ToProj, f)
				if err != nil {
					log.Println("file read:", err)
					return
				}
				variable.Update()
				ChFileModi <- lastEventName
				watchdog++
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}
