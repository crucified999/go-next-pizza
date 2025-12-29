package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-next-pizza/internal/app/model"
	"github.com/go-next-pizza/internal/app/service"
	"github.com/gorilla/mux"
)

type ComboResponse struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Price int `json:"price"`
	Image string `json:"image"`
	Products []*ProductResponse `json:"products"`
	DefaultProducts []*ProductResponse `json:"defaultProducts"`
}

type ComboHandler struct {
	comboService *service.ComboService
}

func NewComboHandler(comboService *service.ComboService) *ComboHandler {
	return &ComboHandler{
		comboService: comboService,
	}
}

func (ch *ComboHandler) GetCombos(w http.ResponseWriter, r *http.Request) {
	combos, err := ch.comboService.GetCombos()
	if err != nil {
		http.Error(w, "Failed to get combos", http.StatusInternalServerError)
		return
	}

	response := make([]*ComboResponse, len(combos))
	for i, combo := range combos {
		response[i] = ch.convertToResponse(combo)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (ch *ComboHandler) GetComboById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	comboId, err := strconv.Atoi(vars["comboId"])
	if err != nil {
		http.Error(w, "Invalid combo ID", http.StatusBadRequest)
		return
	}

	combo, err := ch.comboService.GetComboById(comboId)
	if err != nil {
		http.Error(w, "Failed to get combo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ch.convertToResponse(combo))
}

func (ch *ComboHandler) ReplaceProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	comboId, err := strconv.Atoi(vars["comboId"])

	if err != nil {
		http.Error(w, "Invalid combo ID", http.StatusBadRequest)
		return
	}

	c, err := ch.comboService.GetComboById(comboId)
	if err != nil {
		http.Error(w, "Failed to get combo", http.StatusInternalServerError)
		return
	}

	var req struct {
		ProductToReplaceId int `json:"productToReplaceId"`
		ReplacerId int `json:"replacerId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	combo, err := ch.comboService.ReplaceProduct(req.ProductToReplaceId, req.ReplacerId, c)
	if err != nil {
		http.Error(w, "Failed to replace product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ch.convertToResponse(combo))
}

func (ch *ComboHandler) convertToResponse(combo *model.Combo) *ComboResponse {
	response := &ComboResponse{
		Id: combo.Id,
		Title: combo.Title,
		Description: combo.Description,
		Price: combo.Price,
		Image: combo.Image,
	}

	for _, product := range combo.Products {
		response.Products = append(response.Products, &ProductResponse{
			ID: product.Id,
			Category: product.Category,
			Title: product.Title.String,
			Description: product.Description.String,
			Price: product.Price.Int64,
			Image: product.Image.String,
			Amount: product.Amount.String,
			Weight: product.Weight.Int64,
			Variants: product.Variants,
			Ingredients: product.Ingredients,
			Toppings: product.Toppings,
		})
	}

	for _, product := range combo.DefaultProducts {
		response.DefaultProducts = append(response.DefaultProducts, &ProductResponse{
			ID: product.Id,
			Category: product.Category,
			Title: product.Title.String,
			Description: product.Description.String,
			Price: product.Price.Int64,
			Image: product.Image.String,
			Weight: product.Weight.Int64,
			Amount: product.Amount.String,
			Variants: product.Variants,
			Ingredients: product.Ingredients,
			Toppings: product.Toppings,
		})
	}

	return response
}