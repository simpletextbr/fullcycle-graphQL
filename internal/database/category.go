package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{
		db: db,
	}
}

func (c *Category) Create(name string, descripton string) (Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)",
		id, name, descripton)

	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Name: name, Description: descripton}, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var categories []Category
	for rows.Next() {
		var id, name, description string
		err := rows.Scan(&id, &name, &description)
		if err != nil {
			return nil, err
		}

		categories = append(categories, Category{ID: id, Name: name, Description: description})
	}

	return categories, nil
}

func (c *Category) FindByCourseID(courseId string) (Category, error) {
	var id, name, description string
	err := c.db.QueryRow("SELECT categories.id, categories.name, categories.description FROM categories INNER JOIN courses ON categories.id = courses.category_id WHERE courses.id = $1", courseId).Scan(&id, &name, &description)
	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Name: name, Description: description}, nil
}
