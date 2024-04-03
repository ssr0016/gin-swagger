package main

import (
	"gin-swagger/config"
	"gin-swagger/controller"
	"gin-swagger/helper"
	"gin-swagger/model"
	"gin-swagger/repository"
	"gin-swagger/router"
	"gin-swagger/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Info().Msg("Starting server!")

	// Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	// Repository
	tagsRespository := repository.NewTagsRepositoryImpl(db)

	// Service
	tagsService := service.NewTagsServiceImpl(tagsRespository, validate)

	// Controller
	tagsController := controller.NewTagsController(tagsService)

	// Router
	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrPanic(err)
}
