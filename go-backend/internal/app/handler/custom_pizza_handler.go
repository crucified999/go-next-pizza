package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-next-pizza/internal/app/model"
	"github.com/go-next-pizza/internal/app/service"
	"github.com/gorilla/mux"
)

type CustomPizzaHandler struct {
	customPizzaService *service.CustomPizzaService
}

func NewCustomPizzaHandler(customPizzaService *service.CustomPizzaService) *CustomPizzaHandler {
	return &CustomPizzaHandler{
		customPizzaService: customPizzaService,
	}
}

// CreateCustomPizza создает новую кастомную пиццу
func (cph *CustomPizzaHandler) CreateCustomPizza(w http.ResponseWriter, r *http.Request) {
	var req model.CreateCustomPizzaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value("userID").(int)

	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	customPizza, err := cph.customPizzaService.CreateCustomPizza(userID, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customPizza)
}

func (cph *CustomPizzaHandler) GetCustomPizza(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	customPizza, err := cph.customPizzaService.GetCustomPizzaByID(id)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customPizza)
}

func (cph *CustomPizzaHandler) SetDough(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// userID, ok := r.Context().Value("userID").(int)
	// if !ok {
	// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	// 	return
	// }

	var req model.CustomPizzaDoughRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	customPizza, err := cph.customPizzaService.SetDough(id, req.Dough)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customPizza)
}

// GetCustomPizzas получает все кастомные пиццы пользователя
func (cph *CustomPizzaHandler) GetCustomPizzas(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(int)

	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	customPizzas, err := cph.customPizzaService.GetCustomPizzasByUserID(userID)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customPizzas)
}

// UpdateCustomPizza обновляет кастомную пиццу
func (cph *CustomPizzaHandler) UpdateCustomPizza(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req model.UpdateCustomPizzaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	customPizza, err := cph.customPizzaService.UpdateCustomPizza(id, userID, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customPizza)
}

// DeleteCustomPizza удаляет кастомную пиццу
func (cph *CustomPizzaHandler) DeleteCustomPizza(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Получаем userID из контекста
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if err := cph.customPizzaService.DeleteCustomPizza(id, userID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}