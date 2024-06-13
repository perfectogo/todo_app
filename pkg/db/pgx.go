package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/perfectogo/todo_app/config"
)

func NewPgx(cfg *config.Config) (*pgx.Conn, error) {

	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.DbUser,
		cfg.DbUserPass,
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbName,
	)

	con, err := pgx.Connect(context.Background(), url)
	if err != nil {
		log.Println("error on connecting with database:", err)
		return nil, err
	}

	return con, nil
}
