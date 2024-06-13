package main

import (
	"fmt"
	"log"

	"github.com/perfectogo/todo_app/config"
	"github.com/perfectogo/todo_app/pkg/db"
)

func main() {

	cfg := config.Load()

	con, err := db.NewPgx(cfg)
	if err != nil {
		return
	}
	log.Println("Successfully coneected with database.")

	fmt.Println(con)
}
