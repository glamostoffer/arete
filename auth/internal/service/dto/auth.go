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
	AccessToken string
}
type VerifyCredentialsResponse struct {
	UserID int64
}

type RefreshSessionRequest struct {
	RefreshToken string
}
type RefreshSessionResponse struct {
	AccessToken string
}
