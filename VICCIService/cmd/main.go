package main

import (
	"VICCIService/api"
	"VICCIService/internals/config"
	"VICCIService/internals/entities"
	"VICCIService/internals/repository"
	"VICCIService/internals/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main (){


	configFolder := "./configs/"
	VICCIConfig, err := config.LoadVICCIConfig(configFolder)
	if err != nil {
		log.Fatal(err) // Terminate the application if the config is broken
	}


	newAuth := entities.Auth{
		VICCIConfig.User,
		VICCIConfig.Password,
	}

	router := gin.Default()
	repo := repository.ProvideVICCIRepository(VICCIConfig.Basic_url)
	service := service.ProvideVICCIService(repo)
	api.SetupRouter(router,service, newAuth)

	router.Run()
}