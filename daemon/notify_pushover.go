package zerostick

import (
	"os"

	"github.com/gregdel/pushover"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Pushover configuration

// PushoverClient has the struct for the pushbullet config
type PushoverClient struct {
	UserKey string `json:"user_key"`
	AppKey  string `json:"app_key"`
	Enabled bool   `json:"enabled"`
}

var viperPushoverConfName string = "pushover"

func init() {
	log.Debug("Initializing Pushover")
}

// SaveConfig saves the config
func (po *PushoverClient) SaveConfig() {
	viper.Set(viperPushoverConfName, po)
	viper.WriteConfig()
}

// LoadConfig pulls config from viper
func (po *PushoverClient) LoadConfig() {
	if viper.IsSet(viperPushoverConfName) {
		_ = viper.UnmarshalKey(viperPushoverConfName, po)
	}
}

// SendMessage sends a message via PushBullet
func (po *PushoverClient) SendMessage(message string) {
	app := pushover.New(po.AppKey)
	recipient := pushover.NewRecipient(po.UserKey)
	poMessage := pushover.NewMessage(message)
	response, err := app.SendMessage(poMessage, recipient)
	if err != nil {
		log.Error(err)
	}
	log.Debug(response)
}

// SendMessageWithImage sends a message via PushBullet with a image attached
func (po *PushoverClient) SendMessageWithImage(message, imagePath string) {
	app := pushover.New(po.AppKey)
	recipient := pushover.NewRecipient(po.UserKey)
	poMessage := pushover.NewMessage(message)

	// Open attachment
	file, err := os.Open(imagePath)
	if err != nil {
		log.Error(err)
	}

	// Add attachment
	if err := poMessage.AddAttachment(file); err != nil {
		log.Error(err)
	}

	response, err := app.SendMessage(poMessage, recipient)
	if err != nil {
		log.Error(err)
	}
	log.Debug(response)
}

// DeleteConfig the PushBullet Config
func (po *PushoverClient) DeleteConfig() {
	viper.Set(viperPushbulletConfName, nil)
}
