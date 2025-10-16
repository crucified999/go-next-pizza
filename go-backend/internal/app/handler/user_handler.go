package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-next-pizza/internal/app/storage"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	userRepo storage.UserRepository
}

func NewUserHandler(userRepo storage.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

// func (uh *UserHandler) GetCart(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	userId, err := strconv.Atoi(vars["id"])

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	userCart, err := uh.userRepo.GetCart(userId)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(userCart)
// }

func (uh *UserHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	userOrders, err := uh.userRepo.GetOrders(userId)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userOrders)
}

func (uh *UserHandler) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	user, err := uh.userRepo.FindById(userId)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}