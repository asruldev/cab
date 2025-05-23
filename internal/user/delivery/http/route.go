package http

import (
	"net/http"

	"github.com/asruldev/cab/internal/user/repository"
	"github.com/asruldev/cab/internal/user/usecase"
	"github.com/asruldev/cab/pkg/config"
)

func Route(mux *http.ServeMux) {
	db := config.ConnectPostgres()

	repo := repository.NewPostgresRepo(db)
	uc := usecase.New(repo)
	handler := New(uc)

	mux.HandleFunc("/user", handler.GetUsers)
}
