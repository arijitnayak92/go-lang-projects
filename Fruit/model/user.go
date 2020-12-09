package model

type User struct {
	Email             string
	Password          string
	LoginIP           string
	IsPasswordChanged bool
}

func NewUser(email string, password string, loginIP string, isPasswordChanged bool) *User {
	return &User{Email: email, Password: password, LoginIP: loginIP, IsPasswordChanged: isPasswordChanged}
}
