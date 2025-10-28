package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Env  string `env:"Env" env-default:"dev"`
	Port int    `env:"APP_PORT" env-default:"8080"`
}

func ParseConfig(path string) (*Config, error) {
	cfg := &Config{}
	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
