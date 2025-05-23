package repository

import (
	"github.com/asruldev/cab/internal/user/domain"
	"github.com/jmoiron/sqlx"
)

type PostgresRepo struct {
	db *sqlx.DB
}

func NewPostgresRepo(db *sqlx.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) FindAll() ([]*domain.UserApp, error) {
	var users []*domain.UserApp

	query := "SELECT id, email FROM users"
	err := r.db.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}
