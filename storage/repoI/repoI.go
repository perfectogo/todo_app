package repoi

import (
	"context"

	"github.com/perfectogo/todo_app/models"
)

type UserRepoI interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUsers(ctx context.Context, limit, page int) ([]*models.User, error)
}
