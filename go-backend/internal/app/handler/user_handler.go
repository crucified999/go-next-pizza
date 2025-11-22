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

type ChangeNameRequest struct {
	Name string `json:"name"`
}

type ChangeEmailRequest struct {
	Email string `json:"email"`
}

type ChangeProfileRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

func NewUserHandler(userRepo storage.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

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

func (uh *UserHandler) ChangeName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var req ChangeNameRequest

	json.NewDecoder(r.Body).Decode(&req)

	user, err := uh.userRepo.FindById(userId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := uh.userRepo.ChangeName(user.Id, req.Name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (uh *UserHandler) ChangeEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var req ChangeEmailRequest

	json.NewDecoder(r.Body).Decode(&req)

	user, err := uh.userRepo.FindById(userId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := uh.userRepo.ChangeName(user.Id, req.Email); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// func (uh *UserHandler) ChangeProfile(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	userId, err := strconv.Atoi(vars["id"])

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	var req ChangeProfileRequest

// 	json.NewDecoder(r.Body).Decode(&req)

// 	user, err := uh.userRepo.FindById(userId)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusNotFound)
// 		return
// 	}

// }