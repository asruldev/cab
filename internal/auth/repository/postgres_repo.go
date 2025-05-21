package repository

import (
	"errors"

	"github.com/asruldev/cab/internal/auth/domain"
	"github.com/jmoiron/sqlx"
)

type PostgresRepo struct {
	db *sqlx.DB
}

func NewPostgresRepo(db *sqlx.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) FindByEmail(email string) (*domain.User, error) {
	user := domain.User{}
	query := "SELECT id, email, password FROM users WHERE email = $1"
	err := r.db.Get(&user, query, email)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (r *PostgresRepo) CreateUser(user *domain.User) error {
	query := `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id`
	err := r.db.QueryRow(query, user.Email, user.Password).Scan(&user.ID)
	return err
}
