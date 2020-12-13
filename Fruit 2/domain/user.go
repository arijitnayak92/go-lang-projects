package domain

import (
	"github.com/arijitnayak92/taskAfford/Fruit/model"
)

// GetUser ...
func (d *Domain) GetUser(email string) (*model.User, error) {
	return d.appDB.GetUser(email)
}

// UserSignup ...
func (d *Domain) UserSignup(email string, password string, firstname string, lastname string, role string) (bool, error) {
	return d.appDB.UserSignup(email, password, firstname, lastname, role)

}

// Login ...
func (d *Domain) Login(email string, password string) (string, error) {
	return d.appDB.LoginUser(email, password)

}
