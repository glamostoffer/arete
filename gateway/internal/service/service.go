package service

import (
	auth "github.com/glamostoffer/arete/auth/pkg/api/grpc/v1"
	learning "github.com/glamostoffer/arete/learning/pkg/api/grpc/v1"
)

type service struct {
	auth     auth.AuthClient
	learning learning.LearningClient
}

func New(auth auth.AuthClient, learning learning.LearningClient) *service {
	return &service{
		auth:     auth,
		learning: learning,
	}
}
