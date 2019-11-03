package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/spf13/viper"
	"github.com/zerostick/zerostick/daemon/cmd"
)

const (
	templatesRoot = "./zerostick_web/templates"
	assetsRoot    = "./zerostick_web/assets"
	certsRoot     = "./zerostick_web/certs"
)

func init() {
	viper.SetDefault("templatesRoot", "./zerostick_web/templates")
	viper.SetDefault("assetsRoot", "./zerostick_web/assets")
	viper.SetDefault("certsRoot", "./zerostick_web/certs")

	go func() {
		//Capture program shutdown, to make sure everything shuts down nicely
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		for sig := range c {
			if sig == os.Interrupt {
				log.Print("ZeroStick is shutting down")
				viper.WriteConfig()
				// TODO: Capture and clean
				os.Exit(0)
			}
		}
	}()
}

func main() {
	cmd.Execute()
}
