package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Logger  LoggerConf
	Server  ServerConfig
	Storage StorageConfig
}

type LoggerConf struct {
	Level string
}

type ServerConfig struct {
	Port        string
	Host        string
	Timeout     time.Duration
	IdleTimeout time.Duration
}

type StorageConfig struct {
	path string
}

func NewConfig() (Config, error) {
	var cfg Config

	keys := viper.AllKeys()

	fmt.Println("this keys", keys)

	err := viper.Unmarshal(&cfg)
	if err != nil {
		return cfg, fmt.Errorf("fail to read config %v", err)
	}

	return cfg, nil
}
