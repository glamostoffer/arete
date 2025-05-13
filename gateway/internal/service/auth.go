package service

import (
	"context"

	v1 "github.com/glamostoffer/arete/auth/pkg/api/grpc/v1"
	"github.com/glamostoffer/arete/gateway/internal/service/dto"
)

func (s *service) StartSignUp(ctx context.Context, req dto.StartSignUpRequest) (res dto.StartSignUpResponse, err error) {
	out, err := s.auth.StartSignUp(ctx, &v1.StartSignUpRequest{
		Login:                req.Login,
		Email:                req.Email,
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
	})
	if err != nil {
		return res, err
	}

	return dto.StartSignUpResponse{
		ResendCooldown: out.GetResendCooldown(),
	}, nil
}

func (s *service) ConfirmEmail(ctx context.Context, req dto.ConfirmEmailRequest) (res dto.ConfirmEmailResponse, err error) {
	out, err := s.auth.ConfirmEmail(ctx, &v1.ConfirmEmailRequest{
		ConfirmationCode: req.ConfirmationCode,
		Email:            req.Email,
		Ip:               req.IP,
		UserAgent:        req.UserAgent,
	})
	if err != nil {
		return res, err
	}

	return dto.ConfirmEmailResponse{
		AccessToken:  out.GetAccessToken(),
		RefreshToken: out.GetRefreshToken(),
	}, nil
}

func (s *service) SignIn(ctx context.Context, req dto.SignInRequest) (res dto.SignInResponse, err error) {
	out, err := s.auth.SignIn(ctx, &v1.SignInRequest{
		Login:     req.Login,
		Password:  req.Password,
		Ip:        req.IP,
		UserAgent: req.UserAgent,
	})
	if err != nil {
		return res, err
	}

	return dto.SignInResponse{
		AccessToken:  out.GetAccessToken(),
		RefreshToken: out.GetRefreashToken(),
	}, nil
}

func (s *service) VerifyCredentials(ctx context.Context, req dto.VerifyCredentialsRequest) (res dto.VerifyCredentialsResponse, err error) {
	out, err := s.auth.VerifyCredentials(ctx, &v1.VerifyCredentialsRequest{
		AccessToken: req.AccessToken,
	})
	if err != nil {
		return res, err
	}

	return dto.VerifyCredentialsResponse{
		UserID: out.GetUserID(),
	}, nil
}

func (s *service) RefreshSession(ctx context.Context, req dto.RefreshSessionRequest) (res dto.RefreshSessionResponse, err error) {
	out, err := s.auth.RefreshSession(ctx, &v1.RefreshSessionRequest{
		RefreshToken: req.RefreshToken,
	})
	if err != nil {
		return res, err
	}

	return dto.RefreshSessionResponse{
		AccessToken: out.GetAccessToken(),
	}, nil
}

func (s *service) GetUserInfo(ctx context.Context, req dto.GetUserInfoRequest) (res dto.GetUserInfoResponse, err error) {
	out, err := s.auth.GetUserInfo(ctx, &v1.GetUserInfoRequest{
		UserID: req.UserID,
	})
	if err != nil {
		return res, err
	}

	return dto.GetUserInfoResponse{
		ID:               req.UserID,
		Login:            out.GetLogin(),
		Email:            out.GetEmail(),
		RegistrationDate: out.GetRegistrationDate(),
	}, nil
}
