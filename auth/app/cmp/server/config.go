package server

import "github.com/glamostoffer/arete/pkg/duration"

type Config struct {
	Address           string
	MaxConnectionIdle duration.Duration
	MaxConnectionAge  duration.Duration
	Timeout           duration.Duration
	Time              duration.Duration
	MaxRecvMsgSize    int
	MaxSendMsgSize    int
}
