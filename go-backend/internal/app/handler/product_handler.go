package handler

import (
	"encoding/json"
	"net/http"
	"slices"
	"strconv"

	"github.com/go-next-pizza/internal/app/model"
	"github.com/go-next-pizza/internal/app/service"
	"github.com/gorilla/mux"
)

type ProductResponse struct {
	ID          int     `json:"id"`
	Category    string  `json:"category"`
	Title       string  `json:"title"`
	Description string  `json:"description,omitempty"`
	Price       int64   `json:"price"`
	Image       string  `json:"image"`
	Amount      float64 `json:"amount,omitempty"`
	Weight      int64   `json:"weight"`
}

type ProductHandler struct {
	productService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (ph *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := ph.productService.GetProducts()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := make([]*ProductResponse, len(products))

	for i, product := range products {
		response[i] = ph.convertToResponse(product)
	}

	slices.SortFunc(response, func(a, b *ProductResponse) int {
		return a.ID - b.ID
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (ph *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	product, err := ph.productService.GetProductById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response := ph.convertToResponse(product)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (ph *ProductHandler) GetProdyctsByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]

	products, err := ph.productService.GetProductsByCategory(category)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response := make([]*ProductResponse, len(products))

	for i, product := range products {
		response[i] = ph.convertToResponse(product)
	}

	slices.SortFunc(response, func(a, b *ProductResponse) int {
		return a.ID - b.ID
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (ph *ProductHandler) convertToResponse(product *model.Product) *ProductResponse {
	response := &ProductResponse{
		ID: product.Id,
		Category: product.Category,
	}
	
	if product.Title.Valid {
		response.Title = product.Title.String
	}
	
	if product.Description.Valid {
		response.Description = product.Description.String
	}
	
	if product.Price.Valid {
		response.Price = product.Price.Int64
	}
	
	if product.Image.Valid {
		response.Image = product.Image.String
	}
	
	if product.Amount.Valid {
		response.Amount = product.Amount.Float64
	}
	
	if product.Weight.Valid {
		response.Weight = product.Weight.Int64
	}
	
	return response
}