package models

// CartProduct : Details of individual product added in cart.
type CartProduct struct {
	ID          int
	Quantity    int
	ProductName string
	Price       float64
}

// Cart : Details of Cart.
type Cart struct {
	ID            int
	TotalQuantity int
	Products      []CartProduct
}
