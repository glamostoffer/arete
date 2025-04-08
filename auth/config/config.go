package config

import (
	"encoding/json"
	"os"

	"github.com/glamostoffer/arete/auth/app/cmp/server"
	"github.com/glamostoffer/arete/auth/internal/service"
	"github.com/glamostoffer/arete/auth/pkg/email"
	"github.com/glamostoffer/arete/pkg/psqlconn"
	"github.com/glamostoffer/arete/pkg/redis"
	"github.com/go-playground/validator"
)

type Config struct {
	Service     service.Config
	EmailSender email.Config

	Postgres psqlconn.Config
	Redis    redis.Config
	GRPC     server.Config
}

const (
	path = "./config/config.json"
)

func ReadConfig(cfg *Config) error {
	jsonFile, err := os.Open(path)
	if err != nil {
		return err
	}

	err = json.NewDecoder(jsonFile).Decode(&cfg)
	if err != nil {
		return err
	}

	err = validator.New().Struct(cfg)
	if err != nil {
		return err
	}

	return err
}
