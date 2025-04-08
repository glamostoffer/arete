package domain

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type User struct {
	ID               int64     `json:"id,omitempty" db:"id"`
	Login            string    `json:"login,omitempty" db:"login"`
	Email            string    `json:"email,omitempty" db:"email"`
	HashPassword     string    `json:"-" db:"hash_password"`
	RegistrationDate time.Time `json:"registrationDate,omitempty" db:"registration_date"`
}

func (u *User) GenerateJWT(secret string, ttl time.Duration) (access string, err error) {
	claims := jwt.MapClaims{
		"id":  u.ID,
		"exp": time.Now().Add(ttl).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	access, err = token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return access, nil
}

type UserSession struct {
	User
	IP         string        `json:"ip"`
	UserAgent  string        `json:"userAgent"`
	Expiration time.Duration `json:"-"`
}

type SignUpRequest struct {
	Login        string `json:"login"`
	Email        string `json:"email"`
	HashPassword string `json:"hashPassword"`
	Code         string `json:"code"`
}
