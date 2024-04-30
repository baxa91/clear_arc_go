package main

import (
	"github.com/gin-gonic/gin"
	"omc_go/internal/api"
	"omc_go/internal/repository"
	"omc_go/internal/service"
)

func main() {
	router := gin.Default()
	repo := repository.Repository()
	service := service.Service(repo)
	api.ParseHandler(router, service)
	router.Run(":8080")
}
