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
	Service     service.Config `validate:"required"`
	EmailSender email.Config   `validate:"required"`

	Postgres psqlconn.Config   `validate:"required"`
	Redis    redis.Config      `validate:"required"`
	GRPC     server.ConfigGRPC `validate:"required"`
	HTTP     server.ConfigHTTP `validate:"required"`
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
