package web

import (
	"encoding/json"
	"net/http"

	"github.com/spf13/viper"
	zs "github.com/zerostick/zerostick/daemon"
)

// NabtoConfig returns the Nabto ID
func NabtoConfig(w http.ResponseWriter, r *http.Request) {
	nc := &zs.NabtoClient{}
	if viper.IsSet("nabto") {
		_ = viper.UnmarshalKey("nabto", nc)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(nc)
	w.Write(response)
}

// NabtoSetup takes the Nabto client ID and key and saves the config.
func NabtoSetup(w http.ResponseWriter, r *http.Request) {
	var nc zs.NabtoClient
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&nc)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // "Some problem occurred."
		return
	}
	nc.SetConfig()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(nc)
	w.Write(response)
}

// NabtoDeleteACL handle /nabto/delete_acl
func NabtoDeleteACL(w http.ResponseWriter, r *http.Request) {
	nc := &zs.NabtoClient{}
	nc.DeleteACL()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}
