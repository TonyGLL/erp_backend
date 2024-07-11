package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	Version       string `mapstructure:"VERSION"`
	SMTP_HOST     string `mapstructure:"SMTP_HOST"`
	SMTP_PORT     string `mapstructure:"SMTP_PORT"`
	SMTP_PASSWORD string `mapstructure:"SMTP_PASSWORD"`
	SMTP_FROM     string `mapstructure:"SMTP_FROM"`
}

func LoadConfig(path string, configName string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(configName)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
