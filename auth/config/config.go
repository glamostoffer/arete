package config

import (
	"context"
	"os"

	"github.com/glamostoffer/arete/auth/pkg/api/grpc"
)

type Config struct {
	GRPC grpc.Config `validate:"required"`
}

// todo
func LoadConfig(
	ctx context.Context,
	filePath string,
) (Config, error) {
	_, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}

	return Config{}, nil
}
