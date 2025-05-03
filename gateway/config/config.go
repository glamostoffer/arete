package config

import (
	"encoding/json"
	"os"

	authcli "github.com/glamostoffer/arete/auth/pkg/api/grpc"
	"github.com/glamostoffer/arete/gateway/app/cmp/server"
	learningcli "github.com/glamostoffer/arete/learning/pkg/api/grpc"
	"github.com/go-playground/validator"
)

type Config struct {
	HTTP server.ConfigHTTP `validate:"required"`

	AuthCli     authcli.Config
	LearningCli learningcli.Config
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
