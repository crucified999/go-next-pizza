package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-next-pizza/internal/app/service"
	"github.com/gorilla/mux"
)

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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(combos)
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
	json.NewEncoder(w).Encode(combo)
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
	json.NewEncoder(w).Encode(combo)
}