package service

import (
	GenesisProject "GenesisProject/models"
	"GenesisProject/pkg/repository"
)

type SubscribeService struct {
	repo repository.Subscription
}

func NewSubscribeService(repo repository.Subscription) *SubscribeService {
	return &SubscribeService{repo: repo}
}

func (s *SubscribeService) Subscribe(email GenesisProject.SubscribeEmail) (string, error) {
	return s.repo.Subscribe(email)
}

func (s *SubscribeService) SendEmails() (string, error) {
	return s.repo.SendEmails()
}
