package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zerostick/zerostick/daemon/watchers"
	"github.com/zerostick/zerostick/daemon/web"
)

// serveCmd represents the serve command
var (
	cfgListen = "0.0.0.0:8080"

	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Starte ZeroStick web UI in listening mode",
		Long:  `This will start ZeroStick in web UI mode (only available atm)`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("serve called")
			// Launch TeslaCam watchers
			// Start inotify watcher on the TeslaCam folder
			go func() {
				// Create a gorutine for handling Cam files watcher
				log.Println("Running cam files watcher on", filepath.Join(viper.GetString("cam-root"), "TeslaCam"))
				watchers.CamfilesWatcher(filepath.Join(viper.GetString("cam-root"), "TeslaCam"))
			}()
			go func() {
				watchers.WebfilesWatcher()
			}()
			// Launch web interface
			web.Start()
		},
	}
)

func init() {
	// Add command to root command
	rootCmd.AddCommand(serveCmd)

	// Add local options flags and bind it to viper config
	serveCmd.LocalFlags().StringVarP(&cfgListen, "listen", "l", cfgListen, "hostname:port to listen on")
	viper.BindPFlag("listen", serveCmd.LocalFlags().Lookup("listen"))
}
