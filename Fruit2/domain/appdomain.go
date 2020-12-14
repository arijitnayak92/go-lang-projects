package domain

import (
	"github.com/arijitnayak92/taskAfford/Fruit2/db"
)

// Domain ...
type Domain struct {
	appRepository db.AppRepository
}

// AppDomain ...
type AppDomain interface {
	DatabaseHealthCheck() (error, error)
}

// Product structs
type Product struct {
	productRepo db.ProductRepository
}

// Cart structs
type Cart struct {
	cartRepo db.CartRepository
}

// NewDomain Constructor for Domain struct
func NewDomain(appRepository db.AppRepository) *Domain {
	return &Domain{
		appRepository: appRepository,
	}
}

// ProductDomain Interface
type ProductDomain interface{}

// CartDomain Interface
type CartDomain interface{}

// DatabaseHealthCheck : Fires ping functions of database
func (d *Domain) DatabaseHealthCheck() (error, error) { //(apperrors.ErrPostgresConnection, apperrors.ErrMongoConnection)
	postgresErr := d.appRepository.PostgresHealthCheck()

	mongoErr := d.appRepository.MongoHealthCheck()

	return postgresErr, mongoErr
}
