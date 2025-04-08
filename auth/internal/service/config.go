package service

import (
	"github.com/glamostoffer/arete/pkg/duration"
)

type Config struct {
	ResendCooldown   duration.Duration // todo: make map[attempts]cooldown
	SignUpSessionTTL duration.Duration

	Secret         string
	UserSessionTTL duration.Duration
	AccessTokenTTL duration.Duration
}
