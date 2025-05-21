package server

import (
	"net/http"

	"github.com/asruldev/cab/pkg/config"

	authDelivery "github.com/asruldev/cab/internal/auth/delivery/http"
	authRepo "github.com/asruldev/cab/internal/auth/repository"
	authUsecase "github.com/asruldev/cab/internal/auth/usecase"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Load .env dulu
	config.LoadEnv()

	db := config.ConnectPostgres()

	repo := authRepo.NewPostgresRepo(db)
	uc := authUsecase.New(repo)
	handler := authDelivery.New(uc)

	mux.HandleFunc("/login", handler.Login)
	mux.HandleFunc("/register", handler.Register)
	return mux
}
