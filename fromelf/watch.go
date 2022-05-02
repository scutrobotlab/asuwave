package fromelf

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

var Watcher *fsnotify.Watcher

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
			}
		case err, ok := <-Watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}
