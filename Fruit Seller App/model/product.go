package model

// Product ...
type Product struct {
	ID          int
	Name        string
	Price       int
	ImageID     string
	Description string
}

// NewProduct ...
func NewProduct(ID int, Name string, Price int, ImageID string, Description string) *Product {
	return &Product{ID: ID, Name: Name, Price: Price, ImageID: ImageID, Description: Description}
}
