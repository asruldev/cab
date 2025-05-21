package domain

type User struct {
	ID       string
	Email    string
	Password string // Simpel: plaintext untuk contoh (jangan di prod)
}

type AuthRepository interface {
	FindByEmail(email string) (*User, error)
}

type AuthUsecase interface {
	Login(email, password string) (token string, refreshToken string, err error)
}
