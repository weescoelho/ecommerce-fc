package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/weescoelho/ecommerce-fc/goapi/internal/entity"
	"github.com/weescoelho/ecommerce-fc/goapi/internal/service"
)

type WebProductHandler struct {
	ProductService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *WebProductHandler {
	return &WebProductHandler{ProductService: productService}
}

func (wph *WebProductHandler) GetProducts(writer http.ResponseWriter, request *http.Request) {
	products, err := wph.ProductService.GetProducts()

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(products)
}

func (wch *WebProductHandler) GetProduct(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	if id == "" {
		http.Error(writer, "id is request", http.StatusBadRequest)
	}

	product, err := wch.ProductService.GetProduct(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(product)
}

func (wch *WebProductHandler) GetProductByCategoryId(writer http.ResponseWriter, request *http.Request) {
	categoryId := chi.URLParam(request, "categoryID")
	if categoryId == "" {
		http.Error(writer, "categoryID is request", http.StatusBadRequest)
	}

	products, err := wch.ProductService.GetProductsByCategoryId(categoryId)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(products)
}

func (wch *WebProductHandler) CreateProduct(writer http.ResponseWriter, request *http.Request) {
	var product entity.Product
	err := json.NewDecoder(request.Body).Decode(&product)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := wch.ProductService.CreateProduct(product.Name, product.Description, product.CategoryID, product.ImageURL, product.Price)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(result)
}
