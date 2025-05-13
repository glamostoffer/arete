package server

import (
	"github.com/glamostoffer/arete/pkg/duration"
)

type ConfigHTTP struct {
	Address              string `validate:"required"`
	MaxMultipartMemoryMB int64  `validate:"required"`
	RequestTimeout       duration.Duration
	ReadTimeout          duration.Duration
	WriteTimeout         duration.Duration
	IdleTimeout          duration.Duration
	CORS                 CORSConfig
}

type CORSConfig struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	ExposeHeaders    []string
	AllowCredentials bool
	MaxAge           duration.Duration
}
