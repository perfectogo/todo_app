package storage

import (
	"github.com/jackc/pgx/v5"
	"github.com/perfectogo/todo_app/storage/postgres"
	repoi "github.com/perfectogo/todo_app/storage/repoI"
)

type StorageI interface {
	GetUserRepo() repoi.UserRepoI
}

type storage struct {
	userRepo repoi.UserRepoI
}

func NewStorage(con *pgx.Conn) StorageI {
	return &storage{
		userRepo: postgres.NewUserRepo(con),
	}
}

func (s *storage) GetUserRepo() repoi.UserRepoI {
	return s.userRepo
}
