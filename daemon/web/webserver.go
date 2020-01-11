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
	r = mux.NewRouter().StrictSlash(true) // New Gorilla muxer
}

// LoadTemplates load html templates from file
func LoadTemplates() {
	tpl = template.Must(template.ParseGlob(viper.GetString("templatesRoot") + "/*"))
}

// Start the stuff
func Start() {
	// Load HTML
	LoadTemplates()

	// Scan FS for present TeslaCam recordings in the background
	go func() {
		zshandlers.ScanCamFS(filepath.Join(viper.GetString("cam-root"), "TeslaCam"))
	}()

	// Start inotify watcher on the TeslaCam folder
	//watchers.CamfilesWatcher(filepath.Join(viper.GetString("cam-root"), "TeslaCam"))

	// index
	r.HandleFunc("/", indexPage)
	r.HandleFunc("/index", indexPage)
	r.PathPrefix("/index.html").HandlerFunc(indexPage).Methods("GET") // Show frontpage
	r.Handle("/favicon.ico", http.NotFoundHandler())                  // Return 404 on favicon.ico

	// config
	r.HandleFunc("/config", ConfigPage)
	r.HandleFunc("/post/config", OnPostConfigEvent)

	// events saved
	r.PathPrefix("/events/{type}").HandlerFunc(EventsHandler).Methods("GET")

	// send a video
	r.HandleFunc("/video/{id}", sendVideo).Name("videoRoute")

	// Wifilist scans the network for available SSIDs
	r.PathPrefix("/wifilist").HandlerFunc(Wifilist).Name("wifilist").Methods("GET")
	// Wifi Configuration
	r.PathPrefix("/wifi").HandlerFunc(WifiGetEntries).Name("Wifi Get").Methods("GET")
	r.PathPrefix("/wifi").HandlerFunc(WifiAddEntry).Name("Wifi Add").Methods("POST")
	r.PathPrefix("/wifi/{id}").HandlerFunc(WifiDeleteEntry).Name("Wifi Delete").Methods("DELETE")

	// Nabto
	r.PathPrefix("/nabto").HandlerFunc(NabtoConfig).Name("Nabto Get").Methods("GET")
	r.PathPrefix("/nabto").HandlerFunc(NabtoSetup).Name("Nabto Post").Methods("POST")
	r.PathPrefix("/nabto").HandlerFunc(NabtoDeleteACL).Name("Nabto delete saved creds").Methods("DELETE")

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
