package email

import (
	"context"

	"github.com/glamostoffer/arete/auth/pkg/email/dto"
	"gopkg.in/gomail.v2"
)

type Sender struct {
	cfg Config
}

func New(cfg Config) Sender {
	return Sender{
		cfg: cfg,
	}
}

func (s *Sender) SendHTMLMail(ctx context.Context, req dto.SendEmailRequest) error {
	m := gomail.NewMessage()
	m.SetHeader("To", req.Recipient)
	m.SetHeader("Subject", req.Subject)
	m.SetBody(req.ContentType, req.Body)

	d := gomail.NewDialer(s.cfg.SMTPHost, s.cfg.SMTPPort, s.cfg.Login, s.cfg.Password)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
