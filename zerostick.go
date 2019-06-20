// ZeroStick
// Raspberry Pi Zero W web interface

// +Build ignore
//go:generate go run zerostick_web/assets_generate.go
package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/handlers" // http logging handler
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"gopkg.in/fsnotify.v1"
)

const (
	defaultConfigFile = "./zerostick.yaml"
	templatesRoot     = "./zerostick_web/templates"
	assetsRoot        = "./zerostick_web/assets"
	certsRoot         = "./zerostick_web/certs"
)

var (
	flagConfigFile string
	tpl            *template.Template
)

func init() {
	flag.StringVar(&flagConfigFile, "config", defaultConfigFile, "Sets the path to the configuration file.")
	loadTemplates()

	go func() {
		//Capture program shutdown, to make sure everything shuts down nicely
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		for sig := range c {
			if sig == os.Interrupt {
				log.Print("ZeroStick is shutting down")
				// TODO: Capture and clean
				os.Exit(0)
			}
		}
	}()

	go func() {
		// Create a goroutine for handling webfiles watcher
		webfilesWatcher()
	}()
}

func main() {
	flag.Parse()
	viper.AutomaticEnv()
	viper.SetEnvPrefix("ZEROSTICK") // Preparing for ability to take env variables as input

	log.Println("ZeroStick is starting up")
	log.Printf("Loading configuration from %s\n", flagConfigFile)
	if flagConfigFile == defaultConfigFile {
		log.Println("You can set the location of the config file with the -config flag")
	}
	viper.SetConfigFile(flagConfigFile)

	r := mux.NewRouter() // Gorilla muxer

	r.HandleFunc("/index", index)
	r.HandleFunc("/config", config)
	r.Handle("/favicon.ico", http.NotFoundHandler())
	fs := http.FileServer(http.Dir(assetsRoot))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	log.Print("Listening with TLS on *:10443 (Also https://localhost:10443)")

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServeTLS(":10443", certsRoot+"/cert.pem", certsRoot+"/key.pem", loggedRouter)
}

// Load html templates
func loadTemplates() {
	// log. Println("Reloading templates")
	tpl = template.Must(template.ParseGlob(templatesRoot + "/*"))
}

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
					loadTemplates()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(templatesRoot)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
func config(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "config.gohtml", nil)
}
