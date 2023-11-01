package main

import (
	"fmt"
	"log"
	"net/http"
	"simple-auth/internal/bootstrap"
	"simple-auth/internal/config"
	handlers "simple-auth/internal/handler"
	"simple-auth/internal/repositories"
	"simple-auth/internal/router"
	services "simple-auth/internal/service"
)

func main() {
	config, err := config.ReadConfig("internal/config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
		return
	}

	db, errDb := bootstrap.InitDatabase(config)
	if errDb != nil {
		log.Fatalf("Failed to initialize the database: %v\n", errDb)
		return
	}

	userRepository := repositories.NewUserRepository(db)
	authService := services.NewAuthenticationService(userRepository)
	authHandler := handlers.NewAuthenticationHandler(authService)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	r := router.SetupRouter(authHandler, userHandler)

	port := 8080
	fmt.Printf("Server is running on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
