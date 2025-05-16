package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/glamostoffer/arete/pkg/component"
	"github.com/glamostoffer/arete/pkg/kafka/producer"
	"github.com/glamostoffer/arete/pkg/psqlconn"
	"github.com/glamostoffer/arete/pkg/redis"
	"github.com/glamostoffer/arete/practice/app/cmp/server"
	"github.com/glamostoffer/arete/practice/config"
	"github.com/glamostoffer/arete/practice/internal/cache"
	"github.com/glamostoffer/arete/practice/internal/repository"
	"github.com/glamostoffer/arete/practice/internal/service"
	grpchandler "github.com/glamostoffer/arete/practice/internal/transport/grpc"
)

func Run(ctx context.Context, cfg *config.Config) error {
	psql := psqlconn.New(cfg.Postgres)
	rd := redis.New(cfg.Redis)
	prod := producer.New(cfg.Producer)

	repo := repository.New(psql.DB)
	ch := cache.New(rd.Client)
	// todo producer

	srv := service.New(
		cfg.Service,
		repo,
		ch,
	)

	grpcHandler := grpchandler.New(srv)

	grpcServ := server.NewGRPC(
		cfg.GRPC,
		grpcHandler,
	)

	cmps := []component.Component{
		&psql,
		&rd,
		&prod,
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
