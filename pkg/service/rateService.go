package service

import (
	"GenesisProject/pkg/repository"
)

type RateService struct {
	repo repository.Rate
}

func NewRateService(repo repository.Rate) *RateService {
	return &RateService{repo: repo}
}

func (s *RateService) RateCurrent() (int, error) {
	return s.repo.RateCurrent()
}
