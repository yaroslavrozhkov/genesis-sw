package main

import (
	"GenesisProject"
	"GenesisProject/pkg/handler"
	"GenesisProject/pkg/repository"
	"GenesisProject/pkg/service"
	"log"
)

func main() {
	//handlers := new(handler.Handler)
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(GenesisProject.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error %s", err.Error())
	}
}
