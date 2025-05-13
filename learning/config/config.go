package config

import (
	"encoding/json"
	"os"

	"github.com/glamostoffer/arete/learning/app/cmp/server"
	"github.com/glamostoffer/arete/pkg/psqlconn"
	"github.com/go-playground/validator"
)

type Config struct {
	Postgres psqlconn.Config   `validate:"required"`
	GRPC     server.ConfigGRPC `validate:"required"`
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
