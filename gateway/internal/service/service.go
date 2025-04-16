package service

import (
	v1 "github.com/glamostoffer/arete/auth/pkg/api/grpc/v1"
)

type service struct {
	auth v1.AuthClient
}

func New(auth v1.AuthClient) *service {
	return &service{
		auth: auth,
	}
}
