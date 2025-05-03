package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/glamostoffer/arete/learning/cmp/server"
	"github.com/glamostoffer/arete/learning/config"
	"github.com/glamostoffer/arete/learning/internal/repository"
	"github.com/glamostoffer/arete/learning/internal/service"
	grpchandler "github.com/glamostoffer/arete/learning/internal/transport/grpc"
	"github.com/glamostoffer/arete/pkg/component"
	"github.com/glamostoffer/arete/pkg/psqlconn"
)

func Run(ctx context.Context, cfg *config.Config) error {
	psql := psqlconn.New(cfg.Postgres)

	repo := repository.New(psql.DB)

	srv := service.New(
		repo,
	)

	grpcHandler := grpchandler.New(srv)

	grpcServ := server.NewGRPC(
		cfg.GRPC,
		grpcHandler,
	)

	cmps := []component.Component{
		&psql,
		&grpcServ,
	}

	var err error

	for _, cmp := range cmps { // todo исправить, оптимизировать и вынести
		err = cmp.Start(ctx)
		if err != nil {
			return err
		}
		log.Printf("%s started", cmp.GetName())
	}

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	interruptSignal := <-quitCh

	log.Printf("interrupt signal: %v", interruptSignal)

	for _, cmp := range cmps { // todo исправить, оптимизировать и вынести
		err = cmp.Stop(ctx)
		if err != nil {
			return err
		}

		log.Printf("%s stopped", cmp.GetName())
	}

	return nil
}
