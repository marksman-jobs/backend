package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	MONGO_POOL_MIN             int
	MONGO_POOL_MAX             int
	MONGO_MAX_IDLE_TIME_SECOND int
	MONGO_DATABASE             string
	MONGO_URI                  string
	POSTGRES_URI               string
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
