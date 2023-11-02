package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simple-auth/internal/entity"
	services "simple-auth/internal/service"
)

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) UserHandler {
	return &userHandler{userService}
}

func (h *userHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var response entity.UserResponse

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		response = entity.UserResponse{
			Code:    "01",
			Message: "faield : token missing",
			Data:    nil,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	_, err := h.userService.ValidateToken(tokenString)
	if err != nil {
		response = entity.UserResponse{
			Code:    "01",
			Message: fmt.Sprintf("token validation error : %s", err),
			Data:    nil,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	users, err := h.userService.GetAllUsers()
	if err != nil {
		response = entity.UserResponse{
			Code:    "00",
			Message: fmt.Sprintf("failed : %s", err),
			Data:    users,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	response = entity.UserResponse{
		Code:    "00",
		Message: "success",
		Data:    users,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
