package router

import (
	handlers "simple-auth/internal/handler"

	"github.com/gorilla/mux"
)

func SetupRouter(authHandler handlers.AuthenticationHandler, userHandler handlers.UserHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/login", authHandler.Login).Methods("POST")
	r.HandleFunc("/api/v1/register", authHandler.Register).Methods("POST")
	r.HandleFunc("/api/v1/users", userHandler.GetAllUsers).Methods("GET")
	r.HandleFunc("/api/v1/user", userHandler.GetUserByRefreshToken).Methods("GET")

	return r
}
