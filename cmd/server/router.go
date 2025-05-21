package server

import (
	"net/http"

	_ "github.com/asruldev/cab/docs"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/asruldev/cab/pkg/config"
	"github.com/asruldev/cab/pkg/middleware"
	"github.com/asruldev/cab/pkg/utils"

	authDelivery "github.com/asruldev/cab/internal/auth/delivery/http"
	authRepo "github.com/asruldev/cab/internal/auth/repository"
	authUsecase "github.com/asruldev/cab/internal/auth/usecase"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	db := config.ConnectPostgres()

	repo := authRepo.NewPostgresRepo(db)
	uc := authUsecase.New(repo)
	handler := authDelivery.New(uc)

	mux.HandleFunc("/login", handler.Login)
	mux.HandleFunc("/register", handler.Register)

	// Protected godoc
	// @Summary Protected route
	// @Description Requires valid JWT Bearer token
	// @Tags protected
	// @Security BearerAuth
	// @Success 200 {string} string "Welcome message"
	// @Failure 401 {object} dto.ErrorResponse
	// @Router /protected [get]
	mux.Handle("/protected", middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(middleware.UserContextKey).(*utils.Claims)
		if !ok {
			http.Error(w, "Unauthorized: no user in context", http.StatusUnauthorized)
			return
		}

		w.Write([]byte("Welcome user: " + claims.Email))
	})))

	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	return mux
}
