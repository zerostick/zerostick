package web

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/zerostick/zerostick/daemon/handlers"
)

func sendVideo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars["id"])
	video, err := handlers.CamStructure.FindByID(vars["id"])
	if err != nil { // Not found
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Video was not found")
	}
	log.Println("We got it:", video)
	// Format video for streaming
	tmpVideo, _ := ioutil.TempFile(os.TempDir(), "video_")
	videofile := tmpVideo.Name() + ".mp4"
	handlers.SteamEnableVideo(filepath.Join(video.FullPath, video.Name), videofile)
	http.ServeFile(w, r, videofile)
	//tpl.ExecuteTemplate(w, "sendvideo.gohtml", videofile)
}
