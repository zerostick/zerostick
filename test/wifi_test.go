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

func TestGetWpaSupplicantConf(t *testing.T) {
	wifi := &zs.Wifi{
		SSID:       "flaf",
		Password:   "flaf",
		Priority:   1,
		UseForSync: false,
	}
	wifi.EncryptPassword()
	wifi2 := &zs.Wifi{
		SSID:       "flaf22",
		Password:   "flaf22",
		Priority:   22,
		UseForSync: false,
	}
	wifi2.EncryptPassword()

	wifis := zs.Wifis{}
	wifis.AddWifiToList(*wifi)
	wifis.AddWifiToList(*wifi2)

	expectedConfig := `ctrl_interface=DIR=/var/run/wpa_supplicant GROUP=netdev
update_config=1
country=US
network={
	ssid"flaf"
	psk=21ad66ddf9a61afa2a66d9cf233c722e3993b2dd361b5ca1c3456dd7ea9d8ff4
	priority=1
}
network={
	ssid"flaf22"
	psk=6abd60875676ec4c046945a7773f09ce7b8d49b219a514479a49d50e85b74629
	priority=22
}
`
	generatedConfig := wifis.GetWpaSupplicantConf()
	if generatedConfig != expectedConfig {
		t.Errorf("%s", generatedConfig)
	}
}

func TestWifiList(t *testing.T) {
	scan, err := zs.ScanNetworks()
	if err != nil {
		t.Errorf("ScanNetwork() failed. Check the mocks. %s", err)
	}
	if len(scan) != 3 {
		t.Errorf("Scan of Wifi return not expected: %+v", scan)
	}
}
