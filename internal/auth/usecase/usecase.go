package usecase

import (
	"errors"

	"github.com/asruldev/cab/internal/auth/domain"
	"github.com/asruldev/cab/pkg/utils"
)

// Struct ini menyimpan repo yang merupakan implementasi dari interface AuthRepository. Tujuannya adalah agar AuthUsecase bisa ambil data user dari database melalui kontrak/interface, bukan langsung akses database.
type AuthUsecase struct {
	repo domain.AuthRepository
}

// Fungsi ini adalah constructor: digunakan untuk membuat objek AuthUsecase dan mengisi dependensinya (di sini repo).
// Mengembalikan domain.AuthUsecase, bukan *AuthUsecase, supaya lebih modular dan fleksibel.
func New(repo domain.AuthRepository) domain.AuthUsecase {
	return &AuthUsecase{repo: repo}
}

func (u *AuthUsecase) Login(email, password string) (string, string, error) {
	user, err := u.repo.FindByEmail(email)
	if err != nil {
		return "", "", errors.New("invalid credentials")
	}

	// Cek password pakai bcrypt hash
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", "", errors.New("invalid credentials")
	}

	// Generate JWT access dan refresh token
	token, err := utils.GenerateAccessToken(user.ID, user.Email)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID, user.Email)
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}

func (u *AuthUsecase) Register(email string, password string) (usr string, err error) {
	// Cek apakah user sudah ada
	_, err = u.repo.FindByEmail(email)
	if err == nil {
		return "", errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return "", errors.New("failed to hash password")
	}

	// Buat user baru dengan password sudah di-hash
	newUser := &domain.User{
		Email:    email,
		Password: hashedPassword,
	}

	// Simpan user
	if err := u.repo.CreateUser(newUser); err != nil {
		return "", err
	}

	return newUser.Email, nil
}
