package zerostick

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/xconstruct/go-pushbullet"
)

// Pushbullet configuration

// PushbulletClient has the struct for the pushbullet config
type PushbulletClient struct {
	AccessToken string `json:"access_token"`
	Enabled     bool   `json:"enabled"`
}

var viperPushbulletConfName string = "pushbullet"

func init() {
	log.Debug("Initializing PushBullet")
}

// SaveConfig saves the config
func (pb *PushbulletClient) SaveConfig() {
	viper.Set(viperPushbulletConfName, pb)
	viper.WriteConfig()
}

// LoadConfig pulls config from viper
func (pb *PushbulletClient) LoadConfig() {
	if viper.IsSet(viperPushbulletConfName) {
		_ = viper.UnmarshalKey(viperPushbulletConfName, pb)
	}
}

// SendMessage sends a message via PushBullet
func (pb *PushbulletClient) SendMessage(message string) {
	pbc := pushbullet.New(pb.AccessToken)
	devs, err := pbc.Devices()
	if err != nil {
		log.Warn(err)
	}
	var title string = "" // Empty title shows full message on a device.
	err = pbc.PushNote(devs[0].Iden, title, message)

	if err != nil {
		log.Warn(err)
	}
}

// DeleteConfig the PushBullet Config
func (pb *PushbulletClient) DeleteConfig() {
	viper.Set(viperPushbulletConfName, nil)
}
