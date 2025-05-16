package config

import (
	"encoding/json"
	"os"

	"github.com/glamostoffer/arete/pkg/kafka/producer"
	"github.com/glamostoffer/arete/pkg/psqlconn"
	"github.com/glamostoffer/arete/pkg/redis"
	"github.com/glamostoffer/arete/practice/app/cmp/server"
	"github.com/glamostoffer/arete/practice/internal/service"
	"github.com/go-playground/validator"
)

type Config struct {
	GRPC     server.ConfigGRPC `validate:"required"`
	Service  service.Config    `validate:"required"`
	Postgres psqlconn.Config   `validate:"required"`
	Redis    redis.Config      `validate:"required"`
	Producer producer.Config   `validate:"required"`
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
