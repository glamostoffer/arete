package dto

import "github.com/glamostoffer/arete/auth/internal/domain"

type StartSignUpRequest struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type StartSignUpResponse struct {
	ResendCooldown int64 `json:"resendCooldown"`
}

type ConfirmEmailRequest struct {
	Email            string
	ConfirmationCode string
	IP               string
	UserAgent        string
}
type ConfirmEmailResponse struct {
	AccessToken  string
	RefreshToken string
}

type SignInRequest struct {
	Login     string
	Password  string
	IP        string
	UserAgent string
}
type SignInResponse struct {
	AccessToken  string
	RefreshToken string
}

type GetUserInfoRequest struct {
	UserID int64
}
type GetUserInfoResponse struct {
	domain.User
}
