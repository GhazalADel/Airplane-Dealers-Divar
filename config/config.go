package config

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		PG
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// PG -.
	PG struct {
		HOST     string `env:"POSTGRES_HOST"`
		PORT     string `env:"POSTGRES_PORT"`
		DB       string `env:"POSTGRES_DB"`
		USER     string `env:"POSTGRES_USER"`
		PASSWORD string `env:"POSTGRES_PASSWORD"`
		SSLMODE  string `env:"POSTGRES_SSLMODE"`
		TIMEZONE string `env:"POSTGRES_TIMEZONE"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg, err := ParseConfigFiles("./config/config.yml", "./.env")
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func ParseConfigFiles(files ...string) (*Config, error) {
	var cfg Config

	for i := 0; i < len(files); i++ {
		err := cleanenv.ReadConfig(files[i], &cfg)
		if err != nil {
			log.Printf("Error reading configuration from file:%v", files[i])
			return nil, err
		}
	}

	return &cfg, nil
}
