package service

import (
	"omc_go/internal/domain"
)

type service struct {
	repo domain.Repository
}

func Service(repo domain.Repository) domain.Service {
	return &service{repo: repo}
}

func (s *service) Hello() {
	s.repo.Hello()
	return
}
