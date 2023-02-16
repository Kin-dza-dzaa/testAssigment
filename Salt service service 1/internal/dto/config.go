package dto

import (
	"github.com/spf13/viper"
)

const (
	CONFIG_NAME = "config"
	CONFIG_TYPE = "env"
	CONFIG_PATH = "./config"
)

var Cfg *Config

type Config struct {
	Address string `mapstructure:"ADDRESS"`
}

func ReadConfig() error {
	viper.SetConfigName(CONFIG_NAME)
	viper.SetConfigType(CONFIG_TYPE)
	viper.AddConfigPath(CONFIG_PATH)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&Cfg); err != nil {
		return err
	}
	return nil
}
