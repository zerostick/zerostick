package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile    = "zerostick.yaml"
	cfgCamroot = "/cam"
	cfgDebug   = false

	rootCmd = &cobra.Command{
		Use:   "zerostick",
		Short: "ZeroStick user interface",
		Long:  `ZeroStick user interface and management deamon`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", cfgFile, "config file")
	rootCmd.PersistentFlags().StringVarP(&cfgCamroot, "cam-root", "r", cfgCamroot, "Root of the camera files")
	rootCmd.PersistentFlags().BoolVarP(&cfgDebug, "debug", "d", cfgDebug, "Enable debug output")

	viper.BindPFlag("cam-root", rootCmd.PersistentFlags().Lookup("cam-root"))
	log.Debug(viper.GetString("cam-root"))
}

func initConfig() {

	if cfgDebug {
		log.SetLevel(log.DebugLevel)
		log.SetReportCaller(true)
	}

	viper.SetConfigType("yaml")
	viper.SetConfigFile(cfgFile)
	// Read config
	viper.AutomaticEnv()
	viper.SetEnvPrefix("ZEROSTICK")        // Ie `ZEROSTICK_LISTEN=127.0.0.1:8080 ./zerostick`
	viper.AddConfigPath("/etc/zerostick/") // path to look for the config file in
	viper.AddConfigPath("$HOME/.config/")  // call multiple times to add many search paths
	viper.AddConfigPath(".")               // optionally look for config in the working directory
	log.Info("ZeroStick is starting up")

	err := viper.ReadInConfig()
	if err == nil {
		log.Debugln("Using config file:")
		log.Debugln("Writing additional defaults to the config")
		viper.WriteConfigAs(viper.ConfigFileUsed())
	} else {
		// Ignore errors, we can run without a config file
		log.Infoln("Failed to read in config - But thats OK ;)")
		log.Infoln("Creating a config file with defaults")
		err = viper.WriteConfigAs(viper.ConfigFileUsed())
		if err != nil {
			log.Errorf("Failed to write %s configfile - Error was %v", cfgFile, err)
		}
	}
}

// Execute our daemon
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
