package domain

type UserApp struct {
	ID    string
	Email string
}

type UserRepository interface {
	FindAll() ([]*UserApp, error) // ambil semua user
}

type UserUseCase interface {
	GetAllUsers() ([]*UserApp, error) // ambil semua user
}
