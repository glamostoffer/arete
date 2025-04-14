package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/glamostoffer/arete/auth/app/cmp/server"
	"github.com/glamostoffer/arete/auth/config"
	"github.com/glamostoffer/arete/auth/internal/cache"
	"github.com/glamostoffer/arete/auth/internal/repository"
	"github.com/glamostoffer/arete/auth/internal/service"
	grpchandler "github.com/glamostoffer/arete/auth/internal/transport/grpc"
	httphandler "github.com/glamostoffer/arete/auth/internal/transport/http"
	"github.com/glamostoffer/arete/auth/pkg/email"
	"github.com/glamostoffer/arete/pkg/component"
	"github.com/glamostoffer/arete/pkg/psqlconn"
	"github.com/glamostoffer/arete/pkg/redis"
)

func Run(ctx context.Context, cfg *config.Config) error {
	psql := psqlconn.New(cfg.Postgres)
	rd := redis.New(cfg.Redis)
	sender := email.New(cfg.EmailSender)

	repo := repository.New(psql.DB)
	ch := cache.New(rd.Client)

	srv := service.New(
		cfg.Service,
		sender,
		repo,
		ch,
	)

	grpcHandler := grpchandler.New(srv)
	httpHandler := httphandler.New(srv)

	grpcServ := server.NewGRPC(
		cfg.GRPC,
		grpcHandler,
	)
	httpServ := server.NewHTTP(
		cfg.HTTP,
		httpHandler,
	)

	cmps := []component.Component{
		&psql,
		&rd,
		&grpcServ,
		&httpServ,
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
