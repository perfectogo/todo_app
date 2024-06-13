package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/perfectogo/todo_app/api"
	"github.com/perfectogo/todo_app/api/handlers"
	"github.com/perfectogo/todo_app/config"
	"github.com/perfectogo/todo_app/pkg/db"
	"github.com/perfectogo/todo_app/storage"
)

func main() {

	cfg := config.Load()

	con, err := db.NewPgx(cfg)
	if err != nil {
		return
	}
	log.Println("Successfully coneected with database.")

	storage := storage.NewStorage(con)
	if storage == nil {
		return
	}
	fmt.Println("Successfully got storage.")

	options := api.Options{
		Cfg: cfg,
		HandlerOPtions: handlers.Handlers{
			Storage: storage,
		},
	}

	server := http.Server{
		Addr:    cfg.HttpPort,
		Handler: api.Api(options),
	}

	log.Println("server is ready to server,..")
	log.Println("server is running on", cfg.HttpPort)

	if err := server.ListenAndServe(); err != nil {
		log.Println("error on listening server,")
		return
	}
}
