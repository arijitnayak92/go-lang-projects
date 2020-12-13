package db

import (
	"github.com/arijitnayak92/taskAfford/Fruit/apperrors"
	"github.com/arijitnayak92/taskAfford/Fruit/model"
	"github.com/arijitnayak92/taskAfford/Fruit/utils"
)

// GetUser ...
func (repo *DB) GetUser(email string) (*model.User, error) {
	var t model.User
	query := "SELECT * FROM users WHERE email =$1"
	if err := repo.Postgres.DB.QueryRow(query, email).Scan(&t.Email, &t.Password, &t.FirstName, &t.LastName, &t.Role, &t.CartID); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, apperrors.ErrUserNotFound
		}
		return nil, err
	}
	return &t, nil
}

// UserSignup ...
func (repo *DB) UserSignup(email string, password string, firstname string, lastname string, role string) (bool, error) {
	isPresent, _ := repo.GetUser(email)
	if isPresent != nil {
		return false, apperrors.ErrUserAlreadyPresent
	}
	query := "INSERT INTO users (email, password, firstname, lastname,role,cartid) VALUES ($1, $2, $3, $4,$5,$6)"
	_, err := repo.Postgres.DB.Query(query, email, password, firstname, lastname, role, "")
	if err != nil {
		return false, err
	}
	return false, nil
}

// LoginUser ...
func (repo *DB) LoginUser(email string, password string) (string, error) {
	user, err := repo.GetUser(email)
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
