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
	handler "github.com/glamostoffer/arete/auth/internal/transport/grpc"
	"github.com/glamostoffer/arete/auth/pkg/email"
	"github.com/glamostoffer/arete/pkg/component"
	"github.com/glamostoffer/arete/pkg/psqlconn"
	"github.com/glamostoffer/arete/pkg/redis"
)

func Run(ctx context.Context, cfg *config.Config) error {
	psql := psqlconn.New(cfg.Postgres)
	rd := redis.New(cfg.Redis)
	sender := email.New(cfg.EmailSender)

	log.Printf("psql.DB: %+v", psql.DB)
	repo := repository.New(psql.DB)
	log.Printf("repo: %+v", repo)

	log.Printf("rc.Client: %+v", rd.Client)
	ch := cache.New(rd.Client)
	log.Printf("ch: %+v", ch)

	srv := service.New(
		cfg.Service,
		sender,
		repo,
		ch,
	)

	hdlr := handler.New(srv)

	grpcServ := server.New(
		cfg.GRPC,
		hdlr,
	)

	cmps := []component.Component{
		&psql,
		&rd,
		&grpcServ,
	}

	var err error

	for _, cmp := range cmps { // todo исправить, оптимизировать и вынести
		err = cmp.Start(ctx)
		if err != nil {
			return err
		}
	}

	log.Print("application started!")

	log.Printf("psql.DB: %+v", psql.DB)
	log.Printf("repo: %+v", repo)

	log.Printf("rc.Client: %+v", rd.Client)
	log.Printf("ch: %+v", ch)
	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	interruptSignal := <-quitCh

	log.Printf("interrupt signal: %v", interruptSignal)

	for _, cmp := range cmps { // todo исправить, оптимизировать и вынести
		err = cmp.Stop(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
