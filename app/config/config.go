package config

import (
	"fmt"

	"github.com/joeshaw/envdecode"
)

type Config struct {
	Environment string `env:"ENV,required"`
	Postgres    struct {
		Host     string `env:"POSTGRES_HOST,required"`
		Port     string `env:"POSTGRES_PORT,required"`
		Username string `env:"POSTGRES_USER,required"`
		Password string `env:"POSTGRES_PASSWORD,required"`
		Database string `env:"POSTGRES_DB,required"`
	}
}

// Get returns a config structure.
func Get() Config {
	var cfg Config
	if err := envdecode.Decode(&cfg); err != nil {
		panic(fmt.Sprintf("error to decode config: %s", err))
	}
	return cfg
}
