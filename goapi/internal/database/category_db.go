package database

import (
	"database/sql" // Módulo nátivo do Go para conexão com banco de dados relacionais

	"github.com/weescoelho/ecommerce-fc/goapi/internal/entity"
)

type CategoryDB struct {
	db *sql.DB
}

func NewCategoryDB(db *sql.DB) *CategoryDB {
	return &CategoryDB{db: db}
}

func (cd *CategoryDB) GetCategories() ([]*entity.Category, error) { // Forma de criação de um método
	rows, err := cd.db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close() // defer -> esse comando só será executado quando as demais execuções acontecerem (abaixo)

	var categories []*entity.Category
	for rows.Next() {
		var category entity.Category
		//Itera sobre cada item, retornando para variavel err. Caso err seja vazio (sucesso)
		// Faz o append no slice (array) de categorias

		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}

	return categories, nil
}

func (cd *CategoryDB) GetCategory(id string) (*entity.Category, error) {
	var category entity.Category
	err := cd.db.QueryRow("SELECT id,name FROM categories WHERE id = ?", id).Scan(&category.ID, &category.Name) // .Scan() -> Após a busca faz a atribuição dos dados na variavel (category)

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (cd *CategoryDB) CreateCategory(category *entity.Category) (string, error) {
	_, err := cd.db.Exec("INSERT INTO categories (id,name) VALUES (?, ?)", category.ID, category.Name)
	if err != nil {
		return "", err
	}

	return category.ID, nil
}
