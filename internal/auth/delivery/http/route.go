package http

import (
	"net/http"

	"github.com/asruldev/cab/internal/auth/repository"
	"github.com/asruldev/cab/internal/auth/usecase"
	"github.com/asruldev/cab/pkg/config"
)

func Route(mux *http.ServeMux) {
	db := config.ConnectPostgres()

	repo := repository.NewPostgresRepo(db)
	uc := usecase.New(repo)
	handler := New(uc)

	mux.HandleFunc("/login", handler.Login)
	mux.HandleFunc("/register", handler.Register)
}
