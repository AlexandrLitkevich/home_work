package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Logger        LoggerConf
	Server        ServerConfig
	Storage       Storage
	MemoryStorage MemoryStorage
	SQLStorage    SQLStorage
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

//const (
//	memory = "memory"
//	sql    = "sql"
//)

type Storage struct {
	storageType string
}

type MemoryStorage struct {
	Path string //TODO ???
}

type SQLStorage struct {
	login    string
	password string //TODO how you write???
	Host     string
	Port     string
	Name     string
	Path     string
}

func NewConfig() (Config, error) {
	var cfg Config

	err := viper.Unmarshal(&cfg)
	if err != nil {
		return cfg, fmt.Errorf("fail to read config %v", err)
	}

	return cfg, nil
}
