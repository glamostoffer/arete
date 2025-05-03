package server

import (
	"github.com/glamostoffer/arete/pkg/duration"
)

type ConfigGRPC struct {
	Address           string `validate:"required"`
	MaxConnectionIdle duration.Duration
	MaxConnectionAge  duration.Duration
	Timeout           duration.Duration
	Time              duration.Duration
	MaxRecvMsgSize    int `validate:"required"`
	MaxSendMsgSize    int `validate:"required"`
}
