package server

import (
	"net/http"

	_ "github.com/asruldev/cab/docs"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/asruldev/cab/pkg/middleware"
	"github.com/asruldev/cab/pkg/utils"

	authHttp "github.com/asruldev/cab/internal/auth/delivery/http"
	userHttp "github.com/asruldev/cab/internal/user/delivery/http"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	authHttp.Route(mux)
	userHttp.Route(mux)

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

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "pong"}`))
	}))

	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	return mux
}
