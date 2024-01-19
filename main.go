package main

import (
	"fmt"
	"github.com/saaitt/snappfood_test/app"
	"github.com/spf13/viper"
	"os"
)

func main() {
	fmt.Println(ReadConfig())
	orderManager := app.New(ReadConfig())
	orderManager.Initialize()
	orderManager.Start()
}

func ReadConfig() app.ServiceConfig {
	v := viper.New()
	v.SetConfigType("yaml")
	f, err := os.Open("./config.yaml")
	if err != nil {
		panic("cannot read config: " + err.Error())
	}
	err = v.ReadConfig(f)
	if err != nil {
		panic("cannot read config" + err.Error())
	}
	var configs app.ServiceConfig
	if err := v.Unmarshal(&configs); err != nil {
		fmt.Println(err)
		return configs
	}
	return configs
}
