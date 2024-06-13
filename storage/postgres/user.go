package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/perfectogo/todo_app/models"
	repoi "github.com/perfectogo/todo_app/storage/repoI"
)

type userRepo struct {
	con *pgx.Conn
}

func NewUserRepo(con *pgx.Conn) repoi.UserRepoI {
	return &userRepo{
		con: con,
	}
}

func (u *userRepo) CreateUser(ctx context.Context, user models.User) error {

	query := `
		INSERT INTO 
			users (
				id, 
				username, 
				password, 
				created_at
			) VALUES (
			 	$1, $2, $3, $4
			)`

	_, err := u.con.Exec(
		ctx, query,
		user.ID,
		user.Username,
		user.Password,
		user.CreatedAt,
	)
	if err != nil {
		log.Println("error on insert user to database.")
		return err
	}

	return nil
}

func (u *userRepo) GetUsers(ctx context.Context, limit, page int) ([]*models.User, error) {
	return nil, nil
}
