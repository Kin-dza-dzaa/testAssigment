package dto

import "github.com/spf13/viper"

const (
	CONFIG_PATH = "./config"
	CONFIG_NAME = "config"
	CONFIG_TYPE = "env"
)

var Cfg *Config

type Config struct {
	DbUrl              string `mapstructure:"DB_URL"`
	DbName             string `mapstructure:"DB_NAME"`
	Address            string `mapstructure:"ADDRESS"`
	CollectionName     string `mapstructure:"COLLECTION_NAME"`
	SaltServiceAddress string `mapstructure:"SALT_SERVICE_ADDRESS"`
}

func ReadConfig() error {
	viper.AddConfigPath(CONFIG_PATH)
	viper.SetConfigName(CONFIG_NAME)
	viper.SetConfigType(CONFIG_TYPE)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&Cfg); err != nil {
		return err
	}
	return nil
}
