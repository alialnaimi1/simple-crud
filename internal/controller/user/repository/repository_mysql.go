package repository

import (
	"context"
	"simple-crud/internal/controller/user/data"

	"github.com/jmoiron/sqlx"
)

type MySQLRepository struct {
	*sqlx.DB
}

func NewMySQLRepository(db *sqlx.DB) *MySQLRepository {
	return &MySQLRepository{db}
}

func (r *MySQLRepository) Create(ctx context.Context, username, email, password string) error {
	query := `INSERT INTO users (username, email, password) VALUES (?, ?, ?)`
	_, err := r.ExecContext(ctx, query, username, email, password)
	return err
}

func (r *MySQLRepository) Read(ctx context.Context, id int64) (*data.User, error) {
	query := `SELECT username, email, password FROM users WHERE id = ?`
	var user data.User
	err := r.GetContext(ctx, &user, query, id)
	return &user, err
}

func (r *MySQLRepository) Update(ctx context.Context, username, email, password string) error {
	query := `UPDATE users SET username = ?, email = ?, password = ? WHERE id = ?`
	_, err := r.ExecContext(ctx, query, username, email, password)
	return err
}

func (r *MySQLRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := r.ExecContext(ctx, query, id)
	return err
}
