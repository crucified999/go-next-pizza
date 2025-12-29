package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-next-pizza/internal/app/model"
	"github.com/go-next-pizza/internal/app/storage"
)

type OrderHandler struct {
	orderRepo storage.OrderRepository
}

type CreateOrderRequest struct {
	TotalPrice int `json:"totalPrice"`
}

func NewOrderHandler(orderRepo storage.OrderRepository) *OrderHandler {
	return &OrderHandler{
		orderRepo: orderRepo,
	}
}

func (oh *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	o := &model.Order{}
	o.UserId = r.Context().Value("userID").(int)
	o.CreatedAt = time.Now()

	var req CreateOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	o.TotalPrice = req.TotalPrice

	err := oh.orderRepo.CreateOrder(o)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}