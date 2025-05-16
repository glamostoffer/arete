package service

type service struct {
	cfg      Config
	repo     repository
	cache    cache
	producer producer
}

func New(cfg Config, repo repository, cache cache, producer producer) *service {
	return &service{
		cfg:      cfg,
		repo:     repo,
		cache:    cache,
		producer: producer,
	}
}
