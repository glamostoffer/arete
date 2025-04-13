package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/glamostoffer/arete/auth/internal/domain"
	"github.com/glamostoffer/arete/auth/internal/service/dto"
	"github.com/glamostoffer/arete/auth/pkg/errlist"
	"github.com/golang-jwt/jwt"
)

func (s *service) CreateSession(ctx context.Context, req dto.CreateUserSessionRequest) (res dto.CreateUserSessionResponse, err error) {
	userSession := domain.UserSession{
		User:       req.User,
		IP:         req.IP,
		UserAgent:  req.UserAgent,
		Expiration: s.cfg.UserSessionTTL.Duration,
	}

	accessToken, err := req.User.GenerateJWT(s.cfg.Secret, s.cfg.AccessTokenTTL.Duration)
	if err != nil {
		return res, err
	}

	refresh := make([]byte, 64)
	if _, err := rand.Read(refresh); err != nil {
		return res, err
	}

	refToken := base64.URLEncoding.EncodeToString(refresh)

	err = s.cache.SetUserSession(ctx, refToken, userSession)
	if err != nil {
		return res, err
	}

	return dto.CreateUserSessionResponse{
		AccessToken:  accessToken,
		RefreshToken: refToken,
	}, nil
}

func (s *service) VerifyCredentials(ctx context.Context, req dto.VerifyCredentialsRequest) (res dto.VerifyCredentialsResponse, err error) {
	claims, err := s.validateUserJWT(req.AccessToken, s.cfg.Secret)
	if err != nil {
		return res, err
	}

	userID := (*claims)["id"].(float64)

	return dto.VerifyCredentialsResponse{
		UserID: int64(userID),
	}, nil
}

func (s *service) validateUserJWT(tokenStr, secret string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return nil, errlist.ErrTokenExpired
			}
		} else {
			return nil, errlist.ErrTokenExpired
		}
		return &claims, nil
	}

	return nil, errlist.ErrInvalidToken
}

func (s *service) RefreshSession(ctx context.Context, req dto.RefreshSessionRequest) (res dto.RefreshSessionResponse, err error) {
	session, err := s.cache.GetUserSession(ctx, req.RefreshToken)
	if err != nil {
		return res, err
	}

	accessToken, err := session.User.GenerateJWT(s.cfg.Secret, s.cfg.AccessTokenTTL.Duration)
	if err != nil {
		return res, err
	}

	err = s.cache.SetUserSession(ctx, req.RefreshToken, session)
	if err != nil {
		return res, err
	}

	return dto.RefreshSessionResponse{
		AccessToken: accessToken,
	}, nil
}
