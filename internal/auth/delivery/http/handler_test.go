package http_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	authHttp "github.com/asruldev/cab/internal/auth/delivery/http"
	"github.com/asruldev/cab/internal/auth/dto"
	"github.com/asruldev/cab/internal/auth/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestLogin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Buat mock usecase
	mockUsecase := mocks.NewMockAuthUsecase(ctrl)

	// Inisialisasi handler dengan mock
	handler := authHttp.New(mockUsecase)

	// Data request login
	loginReq := dto.LoginRequest{
		Email:    "test@example.com",
		Password: "secret123",
	}

	// Setup expected behavior
	mockUsecase.
		EXPECT().
		Login(loginReq.Email, loginReq.Password).
		Return("access-token", "refresh-token", nil)

	// Encode request body
	body, _ := json.Marshal(loginReq)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Recorder untuk response
	w := httptest.NewRecorder()

	// Jalankan handler
	handler.Login(w, req)

	// Validasi response
	assert.Equal(t, http.StatusOK, w.Code)

	var resp dto.LoginResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "access-token", resp.Token)
	assert.Equal(t, "refresh-token", resp.RefreshToken)
}
