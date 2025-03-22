package main

import (
	"context"
	"log"

	"github.com/glamostoffer/arete/analytics/config"
	"github.com/glamostoffer/arete/analytics/internal/app"
)

func main() {
	ctx := context.Background()
	cfg := &config.Config{}

	err := config.ReadConfig(cfg)
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	err = app.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
