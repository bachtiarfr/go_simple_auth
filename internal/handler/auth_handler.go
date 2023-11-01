package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	services "simple-auth/internal/service"
)

type authenticationHandler struct {
	authService services.AuthenticationService
}

func NewAuthenticationHandler(authService services.AuthenticationService) AuthenticationHandler {
	return &authenticationHandler{authService}
}

func (h *authenticationHandler) Login(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	email := payload.Email
	password := payload.Password

	token, err := h.authService.Authenticate(email, password)
	if err != nil {
		http.Error(w, fmt.Sprintf("Login failed: %v", err), http.StatusUnauthorized)
		return
	}

	response := struct {
		Code        int    `json:"code"`
		Message     string `json:"message"`
		AccessToken string `json:"access_token"`
	}{
		Code:        00,
		Message:     "success",
		AccessToken: token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
