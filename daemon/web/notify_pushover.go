package web

import (
	"encoding/json"
	"net/http"

	zs "github.com/zerostick/zerostick/daemon"
)

// NotificationPushoverConfig returns the Pushover config
func NotificationPushoverConfig(w http.ResponseWriter, r *http.Request) {
	pb := &zs.PushoverClient{}
	pb.LoadConfig()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(pb)
	w.Write(response)
}

// NotificationPushoverConfigSet takes the Pushover API key and saves the config.
func NotificationPushoverConfigSet(w http.ResponseWriter, r *http.Request) {
	var pb zs.PushoverClient
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&pb)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // "Some problem occurred."
		return
	}
	pb.SaveConfig()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(pb)
	w.Write(response)
}

// NotificationPushoverConfigDelete handle /notifications/provider/pushover
func NotificationPushoverConfigDelete(w http.ResponseWriter, r *http.Request) {
	pb := &zs.PushoverClient{}
	pb.DeleteConfig()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}
