package repository

import GenesisProject "GenesisProject/models"

type Rate interface {
	RateCurrent() (int, error)
}

type Subscription interface {
	Subscribe(email GenesisProject.SubscribeEmail) (string, error)
	SendEmails() (string, error)
}

type Repository struct {
	Rate
	Subscription
}

func NewRepository() *Repository {
	return &Repository{
		Rate:         NewRateRepo(),
		Subscription: NewSubscribeRepo(""),
	}
}
