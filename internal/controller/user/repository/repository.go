package repository

import (
	"context"
	"simple-crud/internal/controller/user/data"
)

type Repository interface {
	Create(ctx context.Context, username, email, password string) error
	Read(ctx context.Context, id int64) (*data.User, error)
	Update(ctx context.Context, username, email, password string) error
	Delete(ctx context.Context, id int64) error
}
