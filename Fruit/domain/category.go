package domain

import (
	"gitlab.com/affordmed/affmed/apperrors"
	"gitlab.com/affordmed/affmed/model"
)

func (d *Domain) CreateCategory(name, description string) (*model.Category, error) {
	var category model.Category

	query := `INSERT INTO "productCategory"(name, description) VALUES ($1, $2) RETURNING id, name, description`
	result := d.pg.QueryRow(query, name, description)

	err := result.Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return nil, apperrors.ErrCreateCategoryFailed
	}

	return &category, nil
}
