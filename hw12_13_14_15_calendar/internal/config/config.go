package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Logger  LoggerConf
	Server  ServerConfig
	Storage Storage
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

type PostgresStore struct {
	Url string
}

type MemoryStore struct {
	Path string
}

type Storage struct {
	StorageType string
	Memory      MemoryStore
	Postgres    PostgresStore
}

func NewConfig() (Config, error) {
	var cfg Config

	err := viper.Unmarshal(&cfg)
	if err != nil {
		return cfg, fmt.Errorf("fail to read config %v", err)
	}

	return cfg, nil
}
