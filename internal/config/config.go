package config

import (
	"github.com/spf13/viper"
)

type (
	Config struct {
		App     App   `mapstructure:"app"`
		Storage MinIO `mapstructure:"minIO"`
	}

	App struct {
		Port        string `mapstructure:"port"`
		DatabaseURL string `mapstructure:"databaseURL"`
	}

	MinIO struct {
		Endpoint          string `mapstructure:"endpoint"`
		AccessKeyId       string `mapstructure:"accessKeyId"`
		SecretAccessKeyId string `mapstructure:"secretAccessKeyId"`
	}
)

func Load(path string) (config *Config, err error) {
	viper.SetConfigFile(path)
	if err = viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err = viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return
}
