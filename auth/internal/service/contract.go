package service

import (
	"context"
	"time"

	"github.com/glamostoffer/arete/auth/internal/domain"
)

type repository interface {
	InsertUser(ctx context.Context, user domain.User) (int64, error)
	GetUserByLoginOrEmail(ctx context.Context, login string) (user domain.User, err error)
	CheckUserExists(ctx context.Context, login, email string) (bool, error)
	UpdateUser(ctx context.Context, user domain.User) error
}

type cache interface {
	SetSignUpRequest(
		ctx context.Context,
		req domain.SignUpRequest,
		ttl time.Duration,
	) error
	GetSignUpRequest(ctx context.Context, sessionID string) (res domain.SignUpRequest, err error)

	SetUserSession(ctx context.Context, refToken string, userSession domain.UserSession) error
	GetUserSession(ctx context.Context, refToken string) (res domain.UserSession, err error)
}
