package service

import (
	models "GenesisProject/models"
	"GenesisProject/pkg/repository"
)

type Rate interface {
	RateCurrent() (int, error)
}

type Subscription interface {
	Subscribe(email models.SubscribeEmail) (string, error)
	SendEmails() (string, error)
}

type Service struct {
	Rate
	Subscription
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Rate:         NewRateService(repos.Rate),
		Subscription: NewSubscribeService(repos.Subscription),
	}
}
