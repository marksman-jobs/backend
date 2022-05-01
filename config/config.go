package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	// TODO: Add config variables here
}

func InitializeConfig() *Config {
	
	fmt.Println("Initialize Config")

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	c := &Config{}

	if err := viper.Unmarshal(c); err != nil {
		panic(fmt.Errorf("fatal error reading config file: %w", err))
	}

	fmt.Println("Config Initialized")

	return c

}
