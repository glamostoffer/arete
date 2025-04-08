package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"

	"github.com/glamostoffer/arete/auth/internal/domain"
	"github.com/glamostoffer/arete/auth/internal/service/dto"
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
	// проверка валидности accessToken

	// продлевать срок жизни refresh токена в redis*

	// возвращать userID из токена
	return res, nil
}

func (s *service) RefreshSession(ctx context.Context, req dto.RefreshSessionRequest) (res dto.RefreshSessionResponse, err error) {
	// проверка валидности refreshToken и поиск его в redis

	// обновление ttl для refresh

	// создание нового accessToken

	// возвращать accessToken
	return res, nil
}
