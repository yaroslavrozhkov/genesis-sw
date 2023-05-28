package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) rate(c *gin.Context) {
	mess, err := h.services.RateCurrent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"description": "Something go wrong",
		})
		return
	}

	c.String(http.StatusOK, "%d", mess)
}
