package watchers

import (
	"github.com/dietsche/rfsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/zerostick/zerostick/daemon/web"
	"gopkg.in/fsnotify.v1"
)

// WebfilesWatcher will monitor templates dir for changes and reload if any
func WebfilesWatcher() {
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
				// log.Println("FS event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("FS modified file:", event.Name)
					web.LoadTemplates()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.AddRecursive(viper.GetString("templatesRoot"))
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
