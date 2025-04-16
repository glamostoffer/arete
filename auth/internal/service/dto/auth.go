package dto

import "github.com/glamostoffer/arete/auth/internal/domain"

type CreateUserSessionRequest struct {
	User      domain.User
	IP        string
	UserAgent string
}
type CreateUserSessionResponse struct {
	AccessToken  string
	RefreshToken string
}

type VerifyCredentialsRequest struct {
	AccessToken string `json:"accessToken"`
}
type VerifyCredentialsResponse struct {
	UserID int64 `json:"userID"`
}

type RefreshSessionRequest struct {
	RefreshToken string
}
type RefreshSessionResponse struct {
	AccessToken string `json:"accessToken"`
}
