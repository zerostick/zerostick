package watchers

import (
	log "github.com/sirupsen/logrus"
	"github.com/zerostick/zerostick/daemon/handlers"
	"gopkg.in/fsnotify.v1"
)

// CamfilesWatcher will monitor folder for new or deleted files.
func CamfilesWatcher(folder string) {
	watcher, err := fsnotify.NewWatcher()
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
					log.Println("FS modified file:", event.Name)
					handlers.HandleCamEvents(event.Name)
				}
				if event.Op&fsnotify.Remove == fsnotify.Create {
					log.Debug("FS Removed file:", event.Name)

				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(folder)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
