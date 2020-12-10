package model

type User struct {
	Email    string
	Password string
}

func NewUser(email string, password string) *User {
	return &User{Email: email, Password: password}
}
