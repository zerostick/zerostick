package web

import (
	"encoding/json"
	"net/http"

	zs "github.com/zerostick/zerostick/daemon"
)

// NotificationPushbulletConfig returns the PushBullet config
func NotificationPushbulletConfig(w http.ResponseWriter, r *http.Request) {
	pb := &zs.PushbulletClient{}
	pb.LoadConfig()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(pb)
	w.Write(response)
}

// NotificationPushbulletConfigSet takes the PushBullet API key and saves the config.
func NotificationPushbulletConfigSet(w http.ResponseWriter, r *http.Request) {
	var pb zs.PushbulletClient
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

// NotificationPushbulletConfigDelete handle /notifications/provider/pushbullet
func NotificationPushbulletConfigDelete(w http.ResponseWriter, r *http.Request) {
	pb := &zs.PushbulletClient{}
	pb.DeleteConfig()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}
