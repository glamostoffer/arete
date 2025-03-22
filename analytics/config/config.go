package config

import (
	"encoding/json"
	"os"

	"github.com/glamostoffer/arete/analytics/internal/eventprocessor"
	"github.com/go-playground/validator"
)

type Config struct {
	EventProcessor eventprocessor.Config
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
