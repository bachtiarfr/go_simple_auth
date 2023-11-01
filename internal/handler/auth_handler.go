package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simple-auth/internal/entity"
	service "simple-auth/internal/service"
)

type authenticationHandler struct {
	authService service.AuthenticationService
}

func NewAuthenticationHandler(authService service.AuthenticationService) AuthenticationHandler {
	return &authenticationHandler{authService}
}

func (h *authenticationHandler) Login(w http.ResponseWriter, r *http.Request) {
	var (
		response entity.LoginResponse
		payload  struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
	)

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	email := payload.Email
	password := payload.Password

	token, err := h.authService.Authenticate(email, password)
	if err != nil {
		response = entity.LoginResponse{
			Code:    "01",
			Message: fmt.Sprintf("failed : %s", err),
			Token:   "",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	response = entity.LoginResponse{
		Code:    "00",
		Message: "success",
		Token:   token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
