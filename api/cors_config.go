package api

import (
	"github.com/caarlos0/env"
)

type corsConfigEnv struct {
	CorsAllowOrigin  string `env:"CORS_ALLOW_ORIGIN" envDefault:"*"`
	CorsAllowMethods string `env:"CORS_ALLOW_METHODS" envDefault:"*"`
}

func NewCorsConfig() (*corsConfig, error) {
	cfg := corsConfigEnv{}
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	return &corsConfig{
		config: cfg,
	}, nil
}

type corsConfig struct {
	config corsConfigEnv
}

func (c *corsConfig) GetAllowOrigin() string {
	return c.config.CorsAllowOrigin
}

func (c *corsConfig) GetAllowMethods() string {
	return c.config.CorsAllowMethods
}
