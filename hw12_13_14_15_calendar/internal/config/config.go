package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Logger      LoggerConf
	Server      ServerConfig
	StorageType string
	SQLStorage  SQLStorage
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

type SQLStorage struct {
	Login    string
	Password string
	Host     string
	Port     string
	DataBase string
	//Path     string
}

func NewConfig() (Config, error) {
	var cfg Config

	err := viper.Unmarshal(&cfg)
	if err != nil {
		return cfg, fmt.Errorf("fail to read config %v", err)
	}

	return cfg, nil
}
