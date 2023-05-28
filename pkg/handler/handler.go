package handler

import (
	"GenesisProject/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.GET("/rate", h.rate)
		api.POST("/subscribe", h.subscribe)
		api.POST("/sendEmails", h.sendEmails)
	}

	return router
}
