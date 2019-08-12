// ZeroStick
// Raspberry Pi Zero W web interface

// +Build ignore
//go:generate go run zerostick_web/assets_generate.go
//go:generate go run zerostick_web/templates_generate.go
package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"time"

	"github.com/gorilla/handlers" // http logging handler
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"gopkg.in/fsnotify.v1"
)

const (
	defaultConfigFile = "zerostick"
	templatesRoot     = "./zerostick_web/templates"
	assetsRoot        = "./zerostick_web/assets"
	certsRoot         = "./zerostick_web/certs"
)

type config struct {
	Port     int
	Hostname string
	// PathMap string `mapstructure:"path_map"`
}

type wifi struct {
	ssid        string
	password    string
	priority    int
	syncEnabled bool
}

// ConfigPageData is exported to use in Config.gohtml
type ConfigPageData struct {
	WifiSsid    string
	HotspotSsid string
}

var (
	flagConfigFile string
	flagHostname   string
	flagPort       int
	tpl            *template.Template
	cfg            config // Config struct
)

func init() {
	flag.StringVar(&flagConfigFile, "configfile", defaultConfigFile, "Sets the path to the configuration file.")
	flag.StringVar(&flagHostname, "hostname", "0.0.0.0", "Hostname to listen on.")
	flag.IntVar(&cfg.Port, "port", 443, "Port number that this program will listen on.")

	loadTemplates()

	go func() {
		//Capture program shutdown, to make sure everything shuts down nicely
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		for sig := range c {
			if sig == os.Interrupt {
				log.Print("ZeroStick is shutting down")
				viper.WriteConfig()
				// TODO: Capture and clean
				os.Exit(0)
			}
		}
	}()

	go func() {
		// Create a goroutine for handling webfiles watcher
		log.Println("Running webfiles watcher...") // TODO: Disable for no dev
		webfilesWatcher()
	}()
}

func main() {
	flag.Parse()

	// Read config
	viper.AutomaticEnv()
	viper.SetEnvPrefix("ZEROSTICK") // Ie ZEROSTICK_PORT=8080 ./zerostick
	log.Println("ZeroStick is starting up")
	log.Printf("Loading configuration from %s\n", flagConfigFile)
	if flagConfigFile == defaultConfigFile {
		log.Println("You can set the location of the config file with the -config flag")
	}
	viper.SetDefault("port", "443")
	viper.SetDefault("hostname", "0.0.0.0")
	viper.AddConfigPath("/etc/zerostick/") // path to look for the config file in
	viper.AddConfigPath("$HOME/.config/")  // call multiple times to add many search paths
	viper.AddConfigPath(".")               // optionally look for config in the working directory
	viper.SetConfigName(flagConfigFile)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		// Ignore errors, we can run without a config file
		log.Println(fmt.Errorf("No config file found: %s", err))
		log.Println("Creating a config file with defaults")
		viper.WriteConfigAs(fmt.Sprintf("./%s.yaml", defaultConfigFile))
	}
	viper.Debug()

	cfg.Port = viper.GetInt("port")
	cfg.Hostname = viper.GetString("hostname")

	viper.WriteConfig() // Write the config to file

	r := mux.NewRouter() // Gorilla muxer

	r.HandleFunc("/", indexPage)
	r.HandleFunc("/index", indexPage)
	r.HandleFunc("/index.html", indexPage)
	r.HandleFunc("/config", configPage)

	r.HandleFunc("/post/config", onPostConfigEvent)
	r.Handle("/favicon.ico", http.NotFoundHandler())

	fs := http.FileServer(http.Dir(assetsRoot))
	//fs := http.FileServer(Assets)
	//var fs http.FileSystem = http.Dir("Assets")
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	log.Print(fmt.Sprintf("Listening with TLS on %s:%d (Maybe even https://localhost:%d)", cfg.Hostname, cfg.Port, cfg.Port))

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	// http server on localhost:8081 for uNabto Tunnel
	srv := &http.Server{
		Handler: loggedRouter,
		Addr:    fmt.Sprintf("%s:%d", "127.0.0.1", 8081),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// https server listening on the configured interface and port
	srvTLS := &http.Server{
		Handler: loggedRouter,
		Addr:    fmt.Sprintf("%s:%d", cfg.Hostname, cfg.Port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	// cert, err := tls.LoadX509KeyPair(certsRoot+"/cert.pem", certsRoot+"/key.pem")
	// if err != nil {
	// 	log.Fatalf("server: loadkeys: %s", err)
	// }
	// srv.TLSConfig = &tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
	go log.Fatal(srv.ListenAndServe()) // Run
	log.Fatal(srvTLS.ListenAndServeTLS(certsRoot+"/cert.pem", certsRoot+"/key.pem"))
	//http.ListenAndServeTLS(fmt.Sprintf("%s:%d", cfg.Hostname, cfg.Port), certsRoot+"/cert.pem", certsRoot+"/key.pem", loggedRouter)
}

// Load html templates
func loadTemplates() {
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

func indexPage(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func configPage(w http.ResponseWriter, r *http.Request) {
	var conf ConfigPageData
	conf.WifiSsid = viper.GetString("wifiSsid")
	conf.HotspotSsid = viper.GetString("hotspotSsid")

	tpl.ExecuteTemplate(w, "config.gohtml", conf)
}

func onPostConfigEvent(w http.ResponseWriter, r *http.Request) {
	ssid := r.FormValue("ssid")
	password := r.FormValue("password")
	formType := r.FormValue("type")

	if formType == "wifi" {
		viper.Set("wifiSsid", ssid)
		viper.Set("wifiPassword", password)
	} else if formType == "hotspot" {
		viper.Set("hotspotSsid", ssid)
		viper.Set("hotspotPassword", password)
	} else {
		http.Error(w, "Unknown type", http.StatusBadRequest)
	}
	viper.WriteConfig()
	// todo: OS level work
}

// This is a function to execute a system command and return output
func getCommandOutput(command string, arguments ...string) string {
	// args... unpacks arguments array into elements
	cmd := exec.Command(command, arguments...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Start()
	if err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
	}
	return out.String()
}
