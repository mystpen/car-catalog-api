package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Host       string `mapstructure:""`
	Port       int
	DBPort     int
	DBName     string
	DBUser     string
	DBPassword string
}

func Load() (*Config, error) {
	config := &Config{}

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return config, nil
}
