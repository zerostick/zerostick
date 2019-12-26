package web

import (
	"crypto/tls"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	zshandlers "github.com/zerostick/zerostick/daemon/handlers"
)

var (
	tpl *template.Template
	r   *mux.Router // Gorilla muxer
)

func init() {
	r = mux.NewRouter() // New Gorilla muxer
}

// LoadTemplates load html templates from file
func LoadTemplates() {
	tpl = template.Must(template.ParseGlob(viper.GetString("templatesRoot") + "/*"))
}

// Start the stuff
func Start() {
	// Load HTML
	LoadTemplates()
	// Scan FS for present TeslaCam recordings
	zshandlers.ScanCamFS(filepath.Join(viper.GetString("cam-root"), "TeslaCam"))

	// Start inotify watcher on the TeslaCam folder
	//watchers.CamfilesWatcher(filepath.Join(viper.GetString("cam-root"), "TeslaCam"))

	// index
	r.HandleFunc("/", indexPage)
	r.HandleFunc("/index", indexPage)
	r.HandleFunc("/index.html", indexPage)
	r.Handle("/favicon.ico", http.NotFoundHandler())

	// config
	r.HandleFunc("/config", ConfigPage)
	r.HandleFunc("/post/config", OnPostConfigEvent)

	// events saved
	r.HandleFunc("/events", eventsSavedPage)

	// send a video
	r.HandleFunc("/video/{id}", sendVideo).Name("videoRoute")

	// Wifi Configuration
	r.HandleFunc("/wifilist", wifilist).Name("wifilist").Methods("GET")
	r.HandleFunc("/wifi/{id}", wifi).Name("wifiRoute").Methods("GET", "POST", "DELETE")

	// Serve assets
	fs := http.FileServer(http.Dir(viper.GetString("assetsRoot")))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	// Serve TeslaCam files
	camFs := http.FileServer(http.Dir(zshandlers.ShadowCamFSPath + "/TeslaCam"))
	r.PathPrefix("/TeslaCam/").Handler(http.StripPrefix("/TeslaCam", camFs))

	log.Infof("Listening with TLS on https://%s (Maybe even http://localhost:8081)", viper.GetString("listen"))

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	// http server on localhost:8081 for uNabto Tunnel
	srv := &http.Server{
		Handler: loggedRouter,
		Addr:    "127.0.0.1:8081",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// https server listening on the configured interface and port
	srvTLS := &http.Server{
		Handler: loggedRouter,
		Addr:    viper.GetString("listen"),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	cert, err := tls.LoadX509KeyPair(filepath.Join(viper.GetString("certsRoot"), "cert.pem"), filepath.Join(viper.GetString("certsRoot"), "key.pem"))
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	srvTLS.TLSConfig = &tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
	go log.Fatal(srv.ListenAndServe()) // Run
	log.Fatal(srvTLS.ListenAndServeTLS(filepath.Join(viper.GetString("certsRoot"), "cert.pem"), filepath.Join(viper.GetString("certsRoot"), "key.pem")))
}
