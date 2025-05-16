package service

import "github.com/glamostoffer/arete/pkg/duration"

type Config struct {
	EventLockTime duration.Duration
	MaxAttempts   int64
}
