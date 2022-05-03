package fromelf

import (
	"log"

	"github.com/fsnotify/fsnotify"
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
	for {
		select {
		case event, ok := <-Watcher.Events:
			if !ok {
				return
			}
			log.Println("event:", event)
			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Println("modified file:", event.Name)
				ChFileModi <- event.Name
			}
		case err, ok := <-Watcher.Errors:
			if !ok {
				return
			}
			ChFileError <- err.Error()
			log.Println("error:", err)
		}
	}
}
