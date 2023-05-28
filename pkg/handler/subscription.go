package handler

import (
	models "GenesisProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) subscribe(c *gin.Context) {
	var input models.SubscribeEmail
	if err := c.BindJSON(&input); err != nil {

	}

	mess, err := h.services.Subscribe(input)
	if err != nil {
		c.JSON(http.StatusConflict, map[string]interface{}{
			"description": mess,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"description": mess,
	})
}

func (h *Handler) sendEmails(c *gin.Context) {
	mess, err := h.services.SendEmails()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"description": "Something go wrong",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"description": mess,
	})
}
