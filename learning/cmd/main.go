package main

import (
	"context"
	"log"

	"github.com/glamostoffer/arete/learning/app"
	"github.com/glamostoffer/arete/learning/config"
)

func main() {
	ctx := context.Background()
	cfg := &config.Config{}

	err := config.ReadConfig(cfg)
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	err = app.Run(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
}
