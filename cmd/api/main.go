package main

import (
	"biami-survey-engine/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/mobile-hook", handler.MobileHookHandler)
	}

	_ = r.Run(":8080")
}
