package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	zs "github.com/zerostick/zerostick/daemon"
)

func Wifilist(w http.ResponseWriter, r *http.Request) {
	// iwlistCmd := exec.Command("iwlist", "wlan0", "scan")
	// iwlistCmdOut, err := iwlistCmd.Output()
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
