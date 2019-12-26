package zerostick

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"

	"golang.org/x/crypto/pbkdf2"
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

// GetWifiConfig returns the Wifi as wpa_supplicant.conf block
func (w Wifi) GetWifiConfig() string {
	return fmt.Sprintf("network={\n\tssid\"%s\"\npsk=\"%s\"\npriority=%d\n}\n", w.SSID, w.EncryptedPassword, w.Priority)
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
