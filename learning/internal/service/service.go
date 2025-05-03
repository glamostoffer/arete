package service

type service struct {
	repo repository
}

func New(repo repository) *service {
	return &service{
		repo: repo,
	}
}
