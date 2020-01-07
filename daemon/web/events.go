package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/zerostick/zerostick/daemon/handlers"
)

func EventsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(handlers.CamStructure.VideoFiles[0].ThumbnailFile, len(handlers.CamStructure.VideoFiles))
	events := handlers.CamStructure.EventsSorted()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(events)
	w.Write(response)
}
