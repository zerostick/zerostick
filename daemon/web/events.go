package web

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zerostick/zerostick/daemon/handlers"
)

// EventsHandler fetches events from the TeslaCam
func EventsHandler(w http.ResponseWriter, r *http.Request) {
	events := handlers.CamStructure.EventsSorted()
	vars := mux.Vars(r)
	typeRequested := vars["type"]
	response, _ := json.Marshal(events[typeRequested])
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
