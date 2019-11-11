package web

import (
	"log"
	"net/http"

	"github.com/zerostick/zerostick/daemon/handlers"
)

func eventsSavedPage(w http.ResponseWriter, r *http.Request) {

	log.Println(handlers.CamStructure.VideoFiles[0].ThumbnailFile, len(handlers.CamStructure.VideoFiles))
	//tpl.ExecuteTemplate(w, "events_saved.gohtml", handlers.CamStructure)
	handlers.CamStructure.EventsSorted()
	tpl.ExecuteTemplate(w, "events_saved.gohtml", handlers.CamStructure.EventsSorted())
}
