package handlers

import "net/http"

type AuthenticationHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
}
