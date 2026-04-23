package handler

import (
	"net/http"

	"biami-engine/internal/model"

	"github.com/gin-gonic/gin"
)

func MobileHookHandler(c *gin.Context) {
	var req model.MobileHookRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	// Placeholder response until session-aware flow is introduced.
	response := "CON Welcome to Biami. \n1. Start Survey \n2. Check Balance"
	c.String(http.StatusOK, response)
}
