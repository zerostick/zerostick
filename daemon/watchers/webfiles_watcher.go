package watchers

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/zerostick/zerostick/daemon/web"
	"gopkg.in/fsnotify.v1"
)

// Monitor templates dir for changes and reload if any
func webfilesWatcher() {
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

	err = watcher.Add(viper.GetString("templatesRoot"))
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
