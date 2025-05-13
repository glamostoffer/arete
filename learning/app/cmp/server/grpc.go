package server

import (
	"context"
	"log"
	"net"

	v1 "github.com/glamostoffer/arete/learning/pkg/api/grpc/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	*grpc.Server

	cfg  ConfigGRPC
	auth v1.LearningServer
}

const (
	grpcComponentName = "grpc-server"
)

func NewGRPC(cfg ConfigGRPC, auth v1.LearningServer) GRPCServer {
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

	v1.RegisterLearningServer(s.Server, s.auth)

	listener, err := net.Listen("tcp", s.cfg.Address)
	if err != nil {
		return err
	}

	go func() {
		if err := s.Server.Serve(listener); err != nil {
			log.Fatalf("GRPC server error: %v", err)
		}
	}()

	return nil
}

func (s *GRPCServer) Stop(ctx context.Context) error {
	s.GracefulStop()
	return nil
}

func (s *GRPCServer) GetName() string {
	return grpcComponentName
}
