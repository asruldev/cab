package http

import (
	"encoding/json"
	"net/http"

	"github.com/asruldev/cab/internal/user/domain"
	"github.com/asruldev/cab/internal/user/dto"
)

type UserHandler struct {
	usecase domain.UserUseCase
}

func New(u domain.UserUseCase) *UserHandler {
	return &UserHandler{usecase: u}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	users, err := h.usecase.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to get users", http.StatusInternalServerError)
		return
	}

	// Konversi domain user ke DTO response
	var resp []dto.UserAppResponse
	for _, u := range users {
		resp = append(resp, dto.UserAppResponse{
			ID:    u.ID,
			Email: u.Email,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
