package grpc

import (
	"context"

	"github.com/glamostoffer/arete/auth/internal/service/dto"
	v1 "github.com/glamostoffer/arete/auth/pkg/api/grpc/v1"
)

type handler struct {
	v1.UnimplementedAuthServer

	service service
}

func New(service service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) StartSignUp(ctx context.Context, req *v1.StartSignUpRequest) (res *v1.StartSignUpResponse, err error) {
	out, err := h.service.StartSignUp(ctx, dto.StartSignUpRequest{
		Login:    req.GetLogin(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return nil, err
	}

	return &v1.StartSignUpResponse{
		ResendCooldown: out.ResendCooldown,
	}, err
}

func (h *handler) ConfirmEmail(ctx context.Context, req *v1.ConfirmEmailRequest) (res *v1.ConfirmEmailResponse, err error) {
	out, err := h.service.ConfirmEmail(ctx, dto.ConfirmEmailRequest{
		Email:            req.GetEmail(),
		ConfirmationCode: req.GetConfirmationCode(),
		IP:               req.GetIp(),
		UserAgent:        req.GetUserAgent(),
	})
	if err != nil {
		return nil, err
	}

	return &v1.ConfirmEmailResponse{
		AccessToken:  out.AccessToken,
		RefreshToken: out.RefreshToken,
	}, err
}

func (h *handler) SignIn(ctx context.Context, req *v1.SignInRequest) (res *v1.SignInResponse, err error) {
	out, err := h.service.SignIn(ctx, dto.SignInRequest{
		Login:     req.GetLogin(),
		Password:  req.GetPassword(),
		IP:        req.GetIp(),
		UserAgent: req.GetUserAgent(),
	})
	if err != nil {
		return nil, err
	}

	return &v1.SignInResponse{
		AccessToken:   out.AccessToken,
		RefreashToken: out.RefreshToken,
	}, nil
}

func (h *handler) VerifyCredentials(ctx context.Context, req *v1.VerifyCredentialsRequest) (res *v1.VerifyCredentialsResponse, err error) {
	out, err := h.service.VerifyCredentials(ctx, dto.VerifyCredentialsRequest{
		AccessToken: req.GetAccessToken(),
	})
	if err != nil {
		return nil, err
	}

	return &v1.VerifyCredentialsResponse{
		UserID: out.UserID,
	}, nil
}

func (h *handler) RefreshSession(ctx context.Context, req *v1.RefreshSessionRequest) (res *v1.RefreshSessionResponse, err error) {
	out, err := h.service.RefreshSession(ctx, dto.RefreshSessionRequest{
		RefreshToken: req.GetRefreshToken(),
	})
	if err != nil {
		return nil, err
	}

	return &v1.RefreshSessionResponse{
		AccessToken: out.AccessToken,
	}, nil
}

func (h *handler) InitiatePasswordChange(ctx context.Context, req *v1.InitiatePasswordChangeRequest) (res *v1.InitiatePasswordChangeResponse, err error) {
	return nil, nil
}

func (h *handler) ConfirmPasswordChange(ctx context.Context, req *v1.ConfirmPasswordChangeRequest) (res *v1.ConfirmPasswordChangeResponse, err error) {
	return nil, nil
}

func (h *handler) ResendEmail(ctx context.Context, req *v1.ResendEmailRequest) (res *v1.ResendEmailResponse, err error) {
	return nil, nil
}

func (h *handler) GetUserInfo(ctx context.Context, req *v1.GetUserInfoRequest) (res *v1.GetUserInfoResponse, err error) {
	out, err := h.service.GetUserInfo(ctx, dto.GetUserInfoRequest{
		UserID: req.GetUserID(),
	})
	if err != nil {
		return nil, err
	}

	return &v1.GetUserInfoResponse{
		Email:            out.Email,
		Login:            out.Login,
		RegistrationDate: out.RegistrationDate.Unix(),
	}, nil
}
