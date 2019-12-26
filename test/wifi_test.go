package main

import (
	"testing"

	zs "github.com/zerostick/zerostick/daemon"
	//_ "github.com/zerostick/zerostick/daemon"
)

func TestWifi(t *testing.T) {
	wifi := &zs.Wifi{
		SSID:       "flaf",
		Password:   "flaf",
		Priority:   1,
		UseForSync: false,
	}
	wifi.EncryptPassword()
	if wifi.Password != "" {
		t.Errorf("Password is not empty after EncryptPassword is called: %s", wifi.Password)
	}
}

func TestWifis(t *testing.T) {
	wifi := &zs.Wifi{
		SSID:       "flaf",
		Password:   "flaf",
		Priority:   1,
		UseForSync: false,
	}
	wifi.EncryptPassword()

	wifis := zs.Wifis{}
	wifis.AddWifiToList(*wifi)
	if len(wifis.Wifis) != 1 {
		t.Errorf("Wifis does not have a single wifi config, but contains %+v", wifis)
	}
}
