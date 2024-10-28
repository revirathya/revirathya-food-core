package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBHostname string `mapstructure:"DB_HOSTNAME"`
	DBPort     int    `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
}

func LoadConfig(path string) (C *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&C)
	return C, err
}
