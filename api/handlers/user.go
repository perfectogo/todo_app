package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/perfectogo/todo_app/models"
)

func (h *handlers) CreateUser(resp http.ResponseWriter, req *http.Request) {
	var user models.User

	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	http.Error(resp, "Error hashing password", http.StatusInternalServerError)
	// 	return
	// }

	user.ID = uuid.New()
	//user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()

	err = h.storage.GetUserRepo().CreateUser(context.Background(), user)
	if err != nil {
		http.Error(resp, "Error creating user", http.StatusInternalServerError)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(user)
}
