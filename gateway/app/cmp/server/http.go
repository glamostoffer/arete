package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	cfg     ConfigHTTP
	handler httpHandler
	eng     *gin.Engine
	srv     *http.Server
}

const (
	httpComponentName = "http-server"
)

func NewHTTP(
	cfg ConfigHTTP,
	handler httpHandler,
) HTTPServer {
	engine := gin.New()

	engine.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	engine.MaxMultipartMemory = cfg.MaxMultipartMemoryMB << 20 // Mb

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      engine,
		ReadTimeout:  cfg.ReadTimeout.Duration,
		WriteTimeout: cfg.WriteTimeout.Duration,
		IdleTimeout:  cfg.IdleTimeout.Duration,
	}

	return HTTPServer{
		cfg:     cfg,
		handler: handler,
		eng:     engine,
		srv:     srv,
	}
}

func (s *HTTPServer) Start(ctx context.Context) error {
	s.setupRoutes()

	go func() {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	return nil
}

func (s *HTTPServer) setupRoutes() {
	if s.cfg.RequestTimeout.Duration > 0 {
		s.eng.Use(func(c *gin.Context) {
			ctx, cancel := context.WithTimeout(c.Request.Context(), s.cfg.RequestTimeout.Duration)
			defer cancel()

			c.Request = c.Request.WithContext(ctx)
			c.Next()
		})
	}

	s.handler.SetupRoutes(s.eng)

	return
}

func (s *HTTPServer) Stop(ctx context.Context) error {
	if s.srv == nil {
		return nil
	}

	shutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(shutCtx); err != nil {
		_ = s.srv.Close()
		return err
	}

	return nil
}

func (s *HTTPServer) GetName() string {
	return httpComponentName
}
