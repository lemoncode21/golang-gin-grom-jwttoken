package main

import (
	"golang-jwttoken/config"
	"golang-jwttoken/controller"
	"golang-jwttoken/helper"
	"golang-jwttoken/model"
	"golang-jwttoken/repository"
	"golang-jwttoken/router"
	"golang-jwttoken/service"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

func main() {

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	//Database
	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()

	db.Table("users").AutoMigrate(&model.Users{})

	//Init Repository
	userRepository := repository.NewUsersRepositoryImpl(db)

	//Init Service
	authenticationService := service.NewAuthenticationServiceImpl(userRepository, validate)

	//Init controller
	authenticationController := controller.NewAuthenticationController(authenticationService)
	usersController := controller.NewUsersController(userRepository)

	routes := router.NewRouter(userRepository, authenticationController, usersController)

	server := &http.Server{
		Addr:           ":" + loadConfig.ServerPort,
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server_err := server.ListenAndServe()
	helper.ErrorPanic(server_err)
}
