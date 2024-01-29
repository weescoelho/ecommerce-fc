package entity // pacote é o nome da pasta

import "github.com/google/uuid"

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"` // Converte para o padrão que deve ser mostrado no JSON
}

func NewCategory(name string) *Category { // * ponteiro -> Resultado alocado em memoria (global)
	return &Category{ //& -> necessário para retornar o ponteiro definido acima
		ID:   uuid.New().String(),
		Name: name,
	}
}

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryID  string  `json:"category_id"`
	ImageURL    string  `json:"image_url"`
}

func NewProduct(name, description string, price float64, categoryID, imageURL string) *Product {
	return &Product{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  categoryID,
		ImageURL:    imageURL,
	}
}
