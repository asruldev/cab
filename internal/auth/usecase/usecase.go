package usecase

import (
	"errors"

	"github.com/asruldev/cab/internal/auth/domain"
)

type AuthUsecase struct {
	repo domain.AuthRepository
}

func New(repo domain.AuthRepository) domain.AuthUsecase {
	return &AuthUsecase{repo: repo}
}

func (u *AuthUsecase) Login(email, password string) (string, string, error) {
	user, err := u.repo.FindByEmail(email)
	if err != nil {
		return "", "", errors.New("invalid credentials")
	}

	// Cek password sederhana (jangan di prod pakai ini)
	if user.Password != password {
		return "", "", errors.New("invalid credentials")
	}

	// Generate token dummy (ganti dengan JWT di produksi)
	token := "dummy_access_token_abc123"
	refreshToken := "dummy_refresh_token_xyz789"

	return token, refreshToken, nil
}

func (u *AuthUsecase) Register(email string, password string) (token string, refreshToken string, err error) {
	// Cek apakah user sudah ada
	_, err = u.repo.FindByEmail(email)
	if err == nil {
		return "", "", errors.New("user already exists")
	}

	// Buat user baru
	newUser := &domain.User{
		Email:    email,
		Password: password, // Note: Jangan plaintext di produksi
	}

	// Simpan user
	if err := u.repo.CreateUser(newUser); err != nil {
		return "", "", err
	}

	// Generate token dummy
	token = "dummy_access_token_abc123"
	refreshToken = "dummy_refresh_token_xyz789"
	return token, refreshToken, nil
}
