package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	zs "github.com/zerostick/zerostick/daemon"
)

// Wifilist scans the network for available SSIDs and returns a list in JSON
func Wifilist(w http.ResponseWriter, r *http.Request) {
	wifiList, err := zs.ScanNetworks()
	if err != nil {
		fmt.Println(err, "Error scanning for networks.")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		jstring, _ := json.Marshal(wifiList)
		fmt.Println(string(jstring))
		w.Write([]byte(jstring))
	}
}
