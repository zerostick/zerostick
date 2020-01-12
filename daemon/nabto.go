package zerostick

import "github.com/spf13/viper"

// Nabto Setup stuff

// NabtoClient has Nabto config
type NabtoClient struct {
	DeviceID  string `json:"deviceid"`
	DeviceKey string `json:"devicekey"`
}

// SetConfig saves the config
func (nc *NabtoClient) SetConfig() {
	viper.Set("nabto", nc)
	viper.WriteConfig()
	nc.ApplyConfig()
}

// ApplyConfig configures systemd to load the variables needed to run the unabto tunnel
func (nc *NabtoClient) ApplyConfig() {

}

// DeleteACL removes the ACL files, so the client need to reconfigure
func (nc *NabtoClient) DeleteACL() {

}
