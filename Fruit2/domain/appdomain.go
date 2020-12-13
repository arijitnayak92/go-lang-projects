package domain

import (
	"gitlab.com/affordmed/fruit-seller-b-backend/db"
)

// Domain struct
type Domain struct {
	//User *User
	// Product       *Product
	// Cart          *Cart
	appRepository db.AppRepository
}

// AppDomain Interface..
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
	// return &Domain{
	// 	User: &User{
	// 		userRepo: appRepository.Postgres,
	// 	},
	// 	Product: &Product{
	// 		productRepo: dbProduct,
	// 	},
	// 	Cart: &Cart{
	// 		cartRepo: dbCart,
	// 	},
	// }
}

// func NewProduct(db db.ProductRepository) *Product {
// 	return &Product{productRepo: db}
// }

// func NewCart(db db.CartRepository) *Cart {
// 	return &Cart{cartRepo: db}
// }

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
