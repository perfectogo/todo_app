package api

import (
	"net/http"

	"github.com/perfectogo/todo_app/api/handlers"
	"github.com/perfectogo/todo_app/config"
)

type Options struct {
	Cfg            *config.Config
	HandlerOPtions handlers.Handlers
}

func Api(options Options) *http.ServeMux {

	h := handlers.NewHadlers(options.HandlerOPtions)

	api := http.NewServeMux()

	// endpoints
	api.HandleFunc("POST /api/create-user", h.CreateUser)

	return api
}
