package http

import (
	"encoding/json"
	"net/http"

	"github.com/asruldev/cab/internal/auth/domain"
	"github.com/asruldev/cab/internal/auth/dto"
)

type AuthHandler struct {
	usecase domain.AuthUsecase
}

func New(u domain.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecase: u}
}

// Login godoc
// @Summary User login
// @Description Authenticate user and return token
// @Tags auth
// @Accept json
// @Produce json
// @Param loginRequest body dto.LoginRequest true "Login data"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} dto.ErrorResponse "Bad request"
// @Failure 401 {object} dto.ErrorResponse "Unauthorized"
// @Router /login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	token, refreshToken, err := h.usecase.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	resp := dto.LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// Register godoc
// @Summary Register new user
// @Description Register user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param registerRequest body dto.RegisterRequest true "Register data"
// @Success 201 {object} dto.RegisterResponse
// @Failure 400 {object} dto.ErrorResponse "Bad request"
// @Router /register [post]
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req dto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	usr, err := h.usecase.Register(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := dto.RegisterResponse{
		User: usr,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
