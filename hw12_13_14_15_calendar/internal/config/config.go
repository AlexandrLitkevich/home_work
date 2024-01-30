package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Logger LoggerConf
	// TODO
}

type LoggerConf struct {
	Level string
	// TODO
}

func NewConfig() (Config, error) {
	var cfg Config

	err := viper.Unmarshal(&cfg)
	if err != nil {
		return cfg, fmt.Errorf("fail to read config %v", err)
	}

	return cfg, nil
}

// TODO
