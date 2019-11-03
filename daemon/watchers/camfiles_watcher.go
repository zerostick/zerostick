package watchers

import (
	"github.com/dietsche/rfsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/zerostick/zerostick/daemon/handlers"
	"gopkg.in/fsnotify.v1"
)

// CamfilesWatcher will monitor folder for new or deleted files.
func CamfilesWatcher(folder string) {
	watcher, err := rfsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Debugln("FS event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Debug("FS modified file:", event.Name)
					handlers.HandleCamEvents(event.Name)
				}
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					log.Debug("FS Removed file:", event.Name)
					handlers.HandleCamEvents(event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.AddRecursive(folder)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
