package handlers

import "github.com/perfectogo/todo_app/storage"

type handlers struct {
	storage storage.StorageI
}

type Handlers struct {
	Storage storage.StorageI
}

func NewHadlers(h Handlers) *handlers {
	return &handlers{
		storage: h.Storage,
	}
}
