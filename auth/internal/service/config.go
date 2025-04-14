package service

import (
	"github.com/glamostoffer/arete/pkg/duration"
)

type Config struct {
	ResendCooldown   duration.Duration `validate:"required"` // todo: make map[attempts]cooldown
	SignUpSessionTTL duration.Duration `validate:"required"`

	Secret         string            `validate:"required"`
	UserSessionTTL duration.Duration `validate:"required"`
	AccessTokenTTL duration.Duration `validate:"required"`
}
