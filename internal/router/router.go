package router

import (
	handlers "simple-auth/internal/handler"

	"github.com/gorilla/mux"
)

func SetupRouter(authHandler handlers.AuthenticationHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/login", authHandler.Login).Methods("POST")

	return r
}
