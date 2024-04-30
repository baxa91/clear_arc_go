package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"omc_go/internal/domain"
)

func ParseHandler(router *gin.Engine, service domain.Service) {
	router.GET("/hello", func(c *gin.Context) {
		service.Hello()
		c.JSON(http.StatusOK, gin.H{"users": "ok"})
	})

	router.NoRoute(func(c *gin.Context) {
		if c.Request.Method != "GET" {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		}
	})
}
