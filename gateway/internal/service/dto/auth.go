package dto

import "time"

type StartSignUpRequest struct {
	Login                string `json:"login"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
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
	ID               int64     `json:"id,omitempty" db:"id"`
	Login            string    `json:"login,omitempty" db:"login"`
	Email            string    `json:"email,omitempty" db:"email"`
	RegistrationDate time.Time `json:"registrationDate,omitempty" db:"registration_date"`
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
