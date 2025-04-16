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
	Email            string `json:"email"`
	ConfirmationCode string `json:"confirmationCode"`
	IP               string `json:"ip"`
	UserAgent        string `json:"userAgent"`
}
type ConfirmEmailResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignInRequest struct {
	Login     string `json:"login"`
	Password  string `json:"password"`
	IP        string `json:"ip"`
	UserAgent string `json:"userAgent"`
}
type SignInResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type GetUserInfoRequest struct {
	UserID int64 `json:"userID" uri:"id"`
}
type GetUserInfoResponse struct {
	domain.User
}
