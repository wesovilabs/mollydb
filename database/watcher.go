package database

import (
	"github.com/fsnotify/fsnotify"
	"log"
	"mollydb/model"
	"mollydb/util"
	"strings"
)

func preWatcher() (*fsnotify.Watcher, chan bool) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("Error %s ", err.Error())
		log.Fatal(err)
	}
	done := make(chan bool)
	return watcher, done
}

//WatchPath function to be invoked
func WatchPath(path string, storage *model.Storage) {
	watcher, done := preWatcher()
	defer watcher.Close()
	go runWatcher(watcher, storage)
	err := watcher.Add(path)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func isSupportedFile(source string) bool {
	return !strings.Contains(source, " ") && util.IsYaml(source)
}

func runWatcher(watcher *fsnotify.Watcher, storage *model.Storage) {
	for {
		select {
		case event := <-watcher.Events:
			if isSupportedFile(event.Name) {
				switch event.Op {
				case fsnotify.Chmod:
					GetInstance().Add(storage, event.Name)
				case fsnotify.Create:
					GetInstance().Add(storage, event.Name)
				case fsnotify.Remove:
					GetInstance().Delete(storage, event.Name)
				}
			}
		case err := <-watcher.Errors:
			log.Println("error:", err)
		}
	}
}
