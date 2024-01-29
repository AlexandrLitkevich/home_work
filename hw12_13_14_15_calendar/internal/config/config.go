package config

type Config struct {
	Logger LoggerConf
	// TODO
}

type LoggerConf struct {
	Level string
	// TODO
}

func NewConfig() Config {
	return Config{}
}

// TODO
