package http

import (
	"context"

	"github.com/glamostoffer/arete/auth/internal/service/dto"
)

type service interface {
	StartSignUp(ctx context.Context, req dto.StartSignUpRequest) (res dto.StartSignUpResponse, err error)
	ConfirmEmail(ctx context.Context, req dto.ConfirmEmailRequest) (res dto.ConfirmEmailResponse, err error)
	SignIn(ctx context.Context, req dto.SignInRequest) (res dto.SignInResponse, err error)

	VerifyCredentials(ctx context.Context, req dto.VerifyCredentialsRequest) (res dto.VerifyCredentialsResponse, err error)
	RefreshSession(ctx context.Context, req dto.RefreshSessionRequest) (res dto.RefreshSessionResponse, err error)

	GetUserInfo(ctx context.Context, req dto.GetUserInfoRequest) (res dto.GetUserInfoResponse, err error)
}
