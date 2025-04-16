package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	authcli "github.com/glamostoffer/arete/auth/pkg/api/grpc"
	"github.com/glamostoffer/arete/gateway/app/cmp/server"
	"github.com/glamostoffer/arete/gateway/config"
	"github.com/glamostoffer/arete/gateway/internal/service"
	httphandler "github.com/glamostoffer/arete/gateway/internal/transport/http"
	"github.com/glamostoffer/arete/pkg/component"
)

func Run(ctx context.Context, cfg *config.Config) error {
	authClient := authcli.New(cfg.AuthCli)

	srv := service.New(authClient)

	httpHandler := httphandler.New(srv)

	httpServ := server.NewHTTP(
		cfg.HTTP,
		httpHandler,
	)

	cmps := []component.Component{
		authClient,
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
