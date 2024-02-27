package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	MessageServiceAddress string `mapstructure:"MESSAGE_SERVICE_ADDRESS" required:"true"`
}

func MustLoadConfig(path string, envName string) (config Config) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(envName)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
		return

	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
		return
	}

	return config
}
