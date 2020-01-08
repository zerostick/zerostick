package web

import (
	"encoding/json"

	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	zs "github.com/zerostick/zerostick/daemon"
)

// WifiGetEntries provides GET to /wifi
func WifiGetEntries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	ws := &zs.Wifis{}
	ws.ParseViperWifi()
	log.Debugf("%+v", ws.List)
	jstr := []byte("{}")
	if ws.List != nil {
		jstr, _ = json.Marshal(ws.List)
	}
	w.Write(jstr)
}

// WifiAddEntry provides POST to /wifi
func WifiAddEntry(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var vifi zs.Wifi
	log.Debugf("request POSTed: %+v", r.Body)
	log.Debugf("vifi POSTed: %+v", vifi)
	err := decoder.Decode(&vifi)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // "Some problem occurred."
		return
	}
	vifi.EncryptPassword()
	ws := &zs.Wifis{}
	ws.ParseViperWifi()
	ws.AddWifiToList(vifi)

	log.Debugf("vifi POSTed: %+v", vifi)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(vifi)
	w.Write(response)
}

// WifiDeleteEntry provides DELETE to /wifi/:id
func WifiDeleteEntry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	ws := &zs.Wifis{}
	ws.ParseViperWifi()
	ws.DeleteWifiFromList(id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// response, _ := json.Marshal(vifi)
	// w.Write(response)
}
