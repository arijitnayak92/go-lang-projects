package model

// User ...
type User struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
	Role      string
	CartID    string
}

// NewUser ...
func NewUser(email string, password string, firstName string, lastName string, role string, cartID string) *User {
	return &User{Email: email, Password: password, FirstName: firstName, LastName: lastName, Role: role, CartID: cartID}
}
