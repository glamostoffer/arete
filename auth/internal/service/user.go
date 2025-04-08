package service

import (
	"context"
	"crypto/rand"
	"math/big"
	"strings"
	"time"

	"github.com/glamostoffer/arete/auth/internal/domain"
	"github.com/glamostoffer/arete/auth/internal/service/dto"
	"github.com/glamostoffer/arete/auth/pkg/email"
	emaildto "github.com/glamostoffer/arete/auth/pkg/email/dto"
	"github.com/glamostoffer/arete/auth/pkg/errlist"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) StartSignUp(ctx context.Context, req dto.StartSignUpRequest) (res dto.StartSignUpResponse, err error) {
	exists, err := s.repo.CheckUserExists(ctx, req.Login, req.Email)
	if err != nil {
		return res, err
	}
	if exists {
		return res, errlist.ErrUserAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return res, err
	}

	code, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		return res, err
	}

	err = s.cache.SetSignUpRequest(
		ctx,
		domain.SignUpRequest{
			Login:        req.Login,
			Email:        req.Email,
			HashPassword: string(hashedPassword),
			Code:         code.String(),
		},
		s.cfg.SignUpSessionTTL.Duration,
	)
	if err != nil {
		return res, err
	}

	err = s.sender.SendHTMLMail(ctx, emaildto.SendEmailRequest{
		Subject:     email.SubjectSignUp,
		Recipient:   req.Email,
		ContentType: email.TypeHTML,
		Body:        strings.Replace(email.TemplateConfirmationCode, "%s", code.String(), 1),
	})
	if err != nil {
		return res, err
	}

	return dto.StartSignUpResponse{
		ResendCooldown: int64(s.cfg.ResendCooldown.Duration),
	}, nil
}

func (s *service) ConfirmEmail(ctx context.Context, req dto.ConfirmEmailRequest) (res dto.ConfirmEmailResponse, err error) {
	signUpReq, err := s.cache.GetSignUpRequest(ctx, req.Email)
	if err != nil {
		return res, err
	}

	if strings.Compare(req.ConfirmationCode, signUpReq.Code) != 0 {
		return res, errlist.ErrInvalidConfirmationCode
	}

	user := domain.User{
		Login:            signUpReq.Login,
		Email:            signUpReq.Email,
		HashPassword:     signUpReq.HashPassword,
		RegistrationDate: time.Now(),
	}

	user.ID, err = s.repo.InsertUser(ctx, user)
	if err != nil {
		return res, err
	}

	tokens, err := s.CreateSession(ctx, dto.CreateUserSessionRequest{
		User:      user,
		IP:        req.IP,
		UserAgent: req.UserAgent,
	})

	return dto.ConfirmEmailResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func (s *service) SignIn(ctx context.Context, req dto.SignInRequest) (res dto.SignInResponse, err error) {

	return res, nil
}
