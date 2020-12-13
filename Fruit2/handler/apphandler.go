package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/affordmed/fruit-seller-b-backend/domain"
)

// AppHandler interface
type AppHandler interface {
	GetAppHealth(c *gin.Context)
}

// Handler ...
type Handler struct {
	appDomain domain.AppDomain
}

// Product : Product Handler struct.
type Product struct {
	domain domain.ProductDomain
}

// Cart : Cart Handler struct.
type Cart struct {
	domain domain.CartDomain
}

// ProductHandler : ProductHandler Interface.
type ProductHandler interface{}

// CartHandler : CartHandler Interface.
type CartHandler interface{}

// // NewProduct : Constructor for Product struct.
// func NewProduct(domain domain.ProductDomain) *Product {
// 	return &Product{domain: domain}
// }

// // NewCart : Constructor for Cart struct.
// func NewCart(domain domain.CartDomain) *Cart {
// 	return &Cart{domain: domain}
// }

// NewHandler ....
func NewHandler(appDomain domain.AppDomain) *Handler {
	return &Handler{
		appDomain: appDomain,
	}
}

// GetAppHealth : Pass Health status of server and database.
func (app *Handler) GetAppHealth(c *gin.Context) {
	postgresErr, mongoErr := app.appDomain.DatabaseHealthCheck()
	if postgresErr != nil && mongoErr == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"postgresIsAlive": false,
			"mongoIsAlive":    true,
			"serverIsAlive":   true,
		})
		return
	}
	if postgresErr == nil && mongoErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"postgresIsAlive": true,
			"mongoIsAlive":    false,
			"serverIsAlive":   true,
		})
		return
	}
	if postgresErr != nil && mongoErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"postgresIsAlive": false,
			"mongoIsAlive":    false,
			"serverIsAlive":   true,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"postgresIsAlive": true,
		"mongoIsAlive":    true,
		"serverIsAlive":   true,
	})
}
