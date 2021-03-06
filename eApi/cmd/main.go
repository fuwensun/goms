package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gomsx/goms/eApi/internal/app"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

func init() {
	fmt.Println("\n---eApi---")
}

func main() {
	parseFlag()

	log.Infof("app init ......")

	viper.Reset()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(cfgpath)
	viper.ReadInConfig()

	app, clean, err := app.InitApp()
	if err != nil {
		panic(err)
	}
	app.Start()

	log.Infof("app start ......")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Infof("get a signal: %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			clean()
			log.Infof("app stop ......")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
