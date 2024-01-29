package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/weescoelho/ecommerce-fc/goapi/internal/entity"
	"github.com/weescoelho/ecommerce-fc/goapi/internal/service"
)

type WebCategoryHandler struct {
	CategoryService *service.CategoryService
}

func NewWebCategoryHandler(categoryService *service.CategoryService) *WebCategoryHandler {
	return &WebCategoryHandler{CategoryService: categoryService}
}

func (wch *WebCategoryHandler) GetCategories(writer http.ResponseWriter, request *http.Request) {
	categories, err := wch.CategoryService.GetCategories()

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(categories)
}

func (wch *WebCategoryHandler) GetCategory(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	if id == "" {
		http.Error(writer, "id is request", http.StatusBadRequest)
	}

	category, err := wch.CategoryService.GetCategory(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(category)
}

func (wch *WebCategoryHandler) CreateCategory(writer http.ResponseWriter, request *http.Request) {
	var category entity.Category
	err := json.NewDecoder(request.Body).Decode(&category)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := wch.CategoryService.CreateCategory(category.Name)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(result)
}
