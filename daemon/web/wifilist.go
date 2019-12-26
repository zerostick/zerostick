package web

import (
	"fmt"
	"net/http"
	"os/exec"
)

func wifilist(w http.ResponseWriter, r *http.Request) {
	iwlistCmd := exec.Command("iwlist", "wlan0", "scan")
	iwlistCmdOut, err := iwlistCmd.Output()
	if err != nil {
		fmt.Println(err, "Error when getting the interface information.")
	} else {
		fmt.Println(string(iwlistCmdOut))
	}
}
