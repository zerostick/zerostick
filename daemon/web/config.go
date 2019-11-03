package web

import (
	"net/http"

	"github.com/spf13/viper"
)

// ConfigPageData is exported to use in Config.gohtml
type ConfigPageData struct {
	WifiSsid    string
	HotspotSsid string
}

func configPage(w http.ResponseWriter, r *http.Request) {
	var conf ConfigPageData
	conf.WifiSsid = viper.GetString("wifiSsid")
	conf.HotspotSsid = viper.GetString("hotspotSsid")

	tpl.ExecuteTemplate(w, "config.gohtml", conf)
}

func onPostConfigEvent(w http.ResponseWriter, r *http.Request) {
	ssid := r.FormValue("ssid")
	password := r.FormValue("password")
	formType := r.FormValue("type")

	if formType == "wifi" {
		viper.Set("wifiSsid", ssid)
		viper.Set("wifiPassword", password)
	} else if formType == "hotspot" {
		viper.Set("hotspotSsid", ssid)
		viper.Set("hotspotPassword", password)
	} else {
		http.Error(w, "Unknown type", http.StatusBadRequest)
	}
	viper.WriteConfig()
	// todo: OS level work
}
