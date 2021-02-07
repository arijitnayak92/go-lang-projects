package domain

import (
	"gitlab.com/affordmed/fruit-seller-a-backend.git/apperrors"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/model"
	"gitlab.com/affordmed/fruit-seller-a-backend.git/utils"
)

// GetUser ...
func (d *Domain) GetUser(email string) (*model.User, error) {

	var t model.User
	query := "SELECT * FROM userdb WHERE email =$1"
	if err := d.pg.QueryRow(query, email).Scan(&t.Email, &t.Password, &t.FirstName, &t.LastName, &t.Role, &t.CartID); err != nil {

		if err.Error() == "sql: no rows in result set" {
			return nil, apperrors.ErrUserNotFound
		}
		return nil, err
	}
	return &t, nil
}

// UserSignup ...
func (d *Domain) UserSignup(email string, password string, firstname string, lastname string, role string, cartid string) (bool, error) {
	isPresent, _ := d.GetUser(email)
	if isPresent != nil {
		return false, apperrors.ErrUserAlreadyPresent
	}
	query := "INSERT INTO userdb (email, password, firstname, lastname,role,cartid) VALUES ($1, $2, $3, $4,$5,$6)"
	_, err := d.pg.Exec(query, email, password, firstname, lastname, role, cartid)
	if err != nil {
		return false, err
	}
	return false, nil
}

// Login ...
func (d *Domain) Login(email string, password string) (string, error) {
	user, err := d.GetUser(email)
	if err != nil {
		return "", apperrors.ErrUserNotFound
	}
	util := utils.NewUtil()
	if !util.CompareHashedPasswords(password, user.Password) {
		return "", apperrors.ErrWrongPassword
	}
	token, errs := util.CreateToken(user.Email)
	if errs != nil {
		return "", errs
	}
	return token.RefreshToken, nil
}
