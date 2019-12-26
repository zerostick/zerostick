package zerostick

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/pbkdf2"
	"io/ioutil"
	"os/exec"
	"strings"
	"time"
)

// Wifi properties struct
type Wifi struct {
	SSID              string `json:"ssid"`
	Password          string
	EncryptedPassword string `json:"encrypted_password"`
	Priority          int    `json:"priority"`
	UseForSync        bool   `json:"use_for_sync"`
}

// Wifis is a slice of Wifi
type Wifis struct {
	Wifis []Wifi `json:"wifis"`
}

// WpaNetwork defines a wifi network to connect to.
type WpaNetwork struct {
	Bssid       string `json:"bssid"`
	Frequency   string `json:"frequency"`
	SignalLevel string `json:"signal_level"`
	Flags       string `json:"flags"`
	Ssid        string `json:"ssid"`
}

// GetWifiConfig returns the Wifi as wpa_supplicant.conf block
func (w Wifi) GetWifiConfig() string {
	return fmt.Sprintf("network={\n\tssid\"%s\"\n\tpsk=%s\n\tpriority=%d\n}\n", w.SSID, w.EncryptedPassword, w.Priority)
}

// AddWifiToList appends the given Wifi to the list
func (ws *Wifis) AddWifiToList(w Wifi) {
	ws.Wifis = append(ws.Wifis, w)
}

// encryptPassword returns password as a WPA2 formatted hash
func (w Wifi) encryptPassword(ssid string, password string) string {
	dk := pbkdf2.Key([]byte(password), []byte(ssid), 4096, 256, sha1.New)
	WPAKey := hex.EncodeToString(dk)[0:64] // First 64 bytes of hex string
	return WPAKey
}

// EncryptPassword encrypts the password in the Wifi struct and removes the unencrypted password
func (w *Wifi) EncryptPassword() {
	w.EncryptedPassword = w.encryptPassword(w.SSID, w.Password)
	w.Password = ""
}

// GetWpaSupplicantConf generates a wpa_supplicant.conf file
func (ws Wifis) GetWpaSupplicantConf() string {
	config := "ctrl_interface=DIR=/var/run/wpa_supplicant GROUP=netdev\nupdate_config=1\ncountry=US\n"
	for _, w := range ws.Wifis {
		config = config + w.GetWifiConfig()
	}
	return config
}

// WriteConfig writes the config to disk
func (ws Wifis) WriteConfig(wpaSupplicantFile string) error {
	if wpaSupplicantFile == "" {
		wpaSupplicantFile = "/etc/wpa_supplicant/wpa_supplicant.conf"
	}
	err := ioutil.WriteFile(wpaSupplicantFile, []byte(ws.GetWpaSupplicantConf()), 0600)
	if err != nil {
		return err
	}
	return nil
}

// ScanNetworks returns a map of WpaNetwork data structures.
func ScanNetworks() (map[string]WpaNetwork, error) {
	wpaNetworks := make(map[string]WpaNetwork, 0)

	scanOut, err := exec.Command("wpa_cli", "-i", "wlan0", "scan").Output()
	if err != nil {
		//log.Fatal(err)
		return wpaNetworks, err
	}
	scanOutClean := strings.TrimSpace(string(scanOut))

	// wait one second for results
	time.Sleep(1 * time.Second)

	if scanOutClean == "OK" {
		log.Debug("OK scan")
		networkListOut, err := exec.Command("wpa_cli", "-i", "wlan0", "scan_results").Output()
		if err != nil {
			//wpa.Log.Fatal(err)
			return wpaNetworks, err
		}

		networkListOutArr := strings.Split(string(networkListOut), "\n")
		for _, netRecord := range networkListOutArr[1:] {
			if !strings.Contains(netRecord, "[WPA2-PSK-CCMP]") {
				continue
			}

			fields := strings.Fields(netRecord)

			if len(fields) > 4 {
				ssid := strings.Join(fields[4:], " ")
				wpaNetworks[ssid] = WpaNetwork{
					Bssid:       fields[0],
					Frequency:   fields[1],
					SignalLevel: fields[2],
					Flags:       fields[3],
					Ssid:        ssid,
				}
			}
		}

	}

	return wpaNetworks, nil
}
