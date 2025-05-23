package usecase

import "github.com/asruldev/cab/internal/user/domain"

type UserUseCase struct {
	repo domain.UserRepository
}

func New(repo domain.UserRepository) domain.UserUseCase {
	return &UserUseCase{repo: repo}
}

// GetAllUsers implements domain.UserUseCase.
func (u *UserUseCase) GetAllUsers() ([]*domain.UserApp, error) {
	users, err := u.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
