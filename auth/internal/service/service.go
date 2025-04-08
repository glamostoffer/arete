package service

import "github.com/glamostoffer/arete/auth/pkg/email"

type service struct {
	cfg    Config
	repo   repository
	cache  cache
	sender email.Sender
}

func New(cfg Config, sender email.Sender, repo repository, cache cache) *service {
	return &service{
		cfg:    cfg,
		sender: sender,
		repo:   repo,
		cache:  cache,
	}
}
