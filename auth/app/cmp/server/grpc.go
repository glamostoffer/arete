package server

import (
	"context"
	"log"
	"net"

	v1 "github.com/glamostoffer/arete/auth/pkg/api/grpc/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	*grpc.Server

	cfg  Config
	auth v1.AuthServer
}

const (
	componentName = "grpc-server"
)

func New(cfg Config, auth v1.AuthServer) GRPCServer {
	return GRPCServer{
		cfg:  cfg,
		auth: auth,
	}
}

func (s *GRPCServer) Start(ctx context.Context) error {
	s.Server = grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: s.cfg.MaxConnectionIdle.Duration,
			MaxConnectionAge:  s.cfg.MaxConnectionAge.Duration,
			Timeout:           s.cfg.Timeout.Duration,
			Time:              s.cfg.Time.Duration,
		}),
		grpc.MaxRecvMsgSize(s.cfg.MaxRecvMsgSize),
		grpc.MaxSendMsgSize(s.cfg.MaxSendMsgSize),
	)

	reflection.Register(s.Server)

	v1.RegisterAuthServer(s.Server, s.auth)

	listener, err := net.Listen("tcp", s.cfg.Address)
	if err != nil {
		return err
	}

	go func() {
		if err := s.Server.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	return nil
}

func (s *GRPCServer) Stop(ctx context.Context) error {
	s.GracefulStop()
	return nil
}

func (s *GRPCServer) GetName() string {
	return componentName
}
