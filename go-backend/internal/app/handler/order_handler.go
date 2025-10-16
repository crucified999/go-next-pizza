package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-next-pizza/internal/app/model"
	"github.com/go-next-pizza/internal/app/storage"
)

type OrderHandler struct {
	orderRepo storage.OrderRepository
}

func NewOrderHandler(orderRepo storage.OrderRepository) *OrderHandler {
	return &OrderHandler{
		orderRepo: orderRepo,
	}
}

func (oh *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	o := &model.Order{}
	o.UserId = r.Context().Value("user_id").(int)

	var cart *model.Cart

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &cart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	order, err := oh.orderRepo.CreateOrder(cart, o)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}