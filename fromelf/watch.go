package fromelf

import (
	"log"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/scutrobotlab/asuwave/variable"
)

var Watcher *fsnotify.Watcher

var ChFileModi chan string = make(chan string, 10)
var ChFileError chan string = make(chan string, 10)

func FileWatch() {
	var err error
	Watcher, err = fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer Watcher.Close()
	watchdog := 0
	lastEventName := ""
	for {
		select {
		case event, ok := <-Watcher.Events:
			if !ok {
				log.Println("Event not ok")
				return
			}
			log.Println("file event:", event)
			if event.Op&fsnotify.Write == fsnotify.Write {
				lastEventName = event.Name
				watchdog = 0
			}
		case err, ok := <-Watcher.Errors:
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
				variable.UpdateVariables()
				ChFileModi <- lastEventName
				watchdog++
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}
