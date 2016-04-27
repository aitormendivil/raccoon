package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func InitializeConfiguration() error {
	viper.SetConfigName("config")        // name of config file (without extension)
	viper.AddConfigPath("/etc/raccoon/") // path to look for the config file in
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Unable to read Configuration: %s\n", err)
		return err
	}

	return nil
}

func ReadConfiguration() error {
	hostIp := viper.GetString("hostIp")
	log.Printf("Host Ip: %s\n", hostIp)

	return nil
}
