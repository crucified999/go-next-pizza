package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-next-pizza/internal/app/model"
	"github.com/go-next-pizza/internal/app/service"
	"github.com/gorilla/mux"
)

type CartHandler struct {
	cartService *service.CartService
	productService *service.ProductService
}

type ProductInCartRequest struct {
	ProductId int `json:"productId"`
	Amount string `json:"amount"`
}

type DeleteProductRequest struct {
	ProductInCartRequest
	Action string `json:"action"`
}

type DeletePizzaRequest struct {
	PizzaInCartRequest
	Action string `json:"action"`
}

type PizzaInCartRequest struct {
	PizzaId int `json:"pizzaId"`
	Dough int `json:"dough"`
	Size string `json:"size"`
	ToppingsMask int `json:"toppingsMask,omitempty"`
}

type ComboInCartRequest struct {
	ComboID int `json:"comboId"`
}

func NewCartHandler(cartService *service.CartService, productService *service.ProductService) *CartHandler {
	return &CartHandler{
		cartService: cartService,
		productService: productService,
	}
}

func (ch *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userID").(int)
	cart, err := ch.cartService.GetCartByUserId(userId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cart)
}

func (ch *CartHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userID").(int)

	var req ProductInCartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := ch.cartService.AddProduct(userId, req.ProductId, req.Amount)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	product, err := ch.productService.GetProductVariant(req.ProductId, req.Amount)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("No such product variant")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (ch *CartHandler) AddPizza(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userID").(int)

	var req PizzaInCartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := ch.cartService.AddPizza(userId, &model.PizzaVariant{
		PizzaId: req.PizzaId,
		Size: req.Size,
		Dough: req.Dough,
	}, req.ToppingsMask)


	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pizza, err := ch.productService.GetPizzaVariant(&model.PizzaVariant{
		PizzaId: req.PizzaId,
		Size: req.Size,
		Dough: req.Dough,
	})

	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pizza.Toppings, err = ch.cartService.GetCartToppings(req.ToppingsMask)

	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pizza)
}

func (ch *CartHandler) AddCombo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var req ComboInCartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	err = ch.cartService.AddCombo(userId, req.ComboID)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Combo added to cart"})
}

func (ch *CartHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userID").(int)

	var req DeleteProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	if req.Action == "delete" {
		err := ch.cartService.DeleteProduct(userId, req.ProductId, req.Amount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if req.Action == "delete-completely" {
		if err := ch.cartService.DeleteProductCompletely(userId, req.ProductId, req.Amount); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted from cart"})
}	

func (ch *CartHandler) DeletePizza(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userID").(int)
	
	log.Print("Удаления пиццы...")

	var req DeletePizzaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	pizza := &model.PizzaVariant{
		PizzaId: req.PizzaId,
		Size: req.Size,
		Dough: req.Dough,
	}

	log.Print("Данные для удаления пиццы: ", req.PizzaId, req.Size, req.Dough, req.ToppingsMask, req.Action)
	
	if req.Action == "delete" {
		err := ch.cartService.DeletePizza(userId, pizza, req.ToppingsMask)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if req.Action == "delete-completely" {
		if err := ch.cartService.DeletePizzaCompletely(userId, pizza, req.ToppingsMask); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted from cart"})
}

func (ch *CartHandler) DeleteCombo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var req ComboInCartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	err = ch.cartService.DeleteCombo(userId, req.ComboID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Combo deleted from cart"})
}

func (ch *CartHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userID").(int)

	cart, err := ch.cartService.GetCartByUserId(userId) 

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = ch.cartService.Refresh(cart.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Cart refreshed"})
}