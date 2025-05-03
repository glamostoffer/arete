package grpc

import (
	v1 "github.com/glamostoffer/arete/learning/pkg/api/grpc/v1"
)

type handler struct {
	v1.UnimplementedLearningServer
	service service
}

func New(service service) *handler {
	return &handler{
		service: service,
	}
}
