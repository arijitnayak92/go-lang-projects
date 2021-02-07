package domain

import (
	"strings"
	"time"

	"gitlab.com/affordmed/fruit-seller-a-backend.git/apperrors"
)

// CreateProduct ...
func (d *Domain) CreateProduct(name string, price int, imageID string, description string) (int, error) {
	var (
		id int
	)
	query := "INSERT INTO products (id,name,price,imageid,description,createdat,updatedat) VALUES(nextval('prod_id'),$2,$3,$4,$5,$6,$7) RETURNING id;"

	err := d.pg.QueryRow(query, name, price, imageID, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return 0, apperrors.ErrDuplicateEnrty
		}

		return 0, err
	}

	return id, nil

}
