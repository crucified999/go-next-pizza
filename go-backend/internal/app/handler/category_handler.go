package handler

import (
	"encoding/json"
	"net/http"
	"sort"

	"github.com/go-next-pizza/internal/app/storage"
)

type CategoryHandler struct {
	categoryService storage.CategoryRepository
}

func NewCategoryHandler(categoryService storage.CategoryRepository) *CategoryHandler {
	return &CategoryHandler{categoryService: categoryService}
}

func (ch *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := ch.categoryService.GetCategories()
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sort.Slice(categories, func(i, j int) bool {
		return categories[i].ID < categories[j].ID
	})
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
}