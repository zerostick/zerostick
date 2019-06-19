// ZeroStick
// Raspberry Pi Zero W web interface

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
)

const defaultConfigFile = "./config.yaml"

var (
	flagConfigFile string
	tpl            *template.Template
)

func init() {
	flag.StringVar(&flagConfigFile, "config", defaultConfigFile, "Sets the path to the configuration file.")
	tpl = template.Must(template.ParseGlob("zerostick_web/templates/*"))

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

	r.HandleFunc("/", index)
	r.Handle("/favicon.ico", http.NotFoundHandler())
	fs := http.FileServer(http.Dir("./zerostick_web/assets"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	log.Print("Listening with TLS on *:10443 (Also https://localhost:10443)")

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServeTLS(":10443", "zerostick_web/certs/cert.pem", "zerostick_web/certs/key.pem", loggedRouter)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
