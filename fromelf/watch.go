package fromelf

import (
	"log"
	"os"

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
	for {
		select {
		case event, ok := <-Watcher.Events:
			if !ok {
				return
			}
			log.Println("event:", event)
			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Println("modified file:", event.Name)

				file, err := os.Open(event.Name)
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
				ChFileModi <- event.Name
			}
		case err, ok := <-Watcher.Errors:
			if !ok {
				return
			}
			ChFileError <- err.Error()
			log.Println("error:", err)
		default:
			watchdog++
		}
	}
}
