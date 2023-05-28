package repository

import (
	GenesisProject "GenesisProject/models"
	"fmt"
	"strings"
)

type SubscribeRepo struct {
	storage string
}

func NewSubscribeRepo(storage string) *SubscribeRepo {
	return &SubscribeRepo{storage: storage}
}

func (r *SubscribeRepo) Subscribe(email GenesisProject.SubscribeEmail) (string, error) {
	mess, err := checkAndWrite(email.Email)
	return mess, err
}

func (r *SubscribeRepo) SendEmails() (string, error) {
	emails := getEmailsList()

	for _, email := range emails {

		fmt.Println("Відправка листа на пошту : ", strings.Trim(email, "\n"))
	}

	return "E-mailʼи відправлено", nil
}
