package grpc

import (
	"context"

	v1 "github.com/glamostoffer/arete/auth/pkg/api/grpc/v1"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type connector struct {
	v1.AuthClient

	cfg  Config
	conn *grpc.ClientConn
}

func (c *connector) Start(ctx context.Context) error {
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	}

	conn, err := grpc.NewClient(
		c.cfg.Address,
		options...,
	)
	if err != nil {
		return err
	}

	c.conn = conn
	c.AuthClient = v1.NewAuthClient(c.conn)

	return nil
}
