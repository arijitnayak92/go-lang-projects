package domain

import (
	"gitlab.com/affordmed/affmed/apperrors"
	"gitlab.com/affordmed/affmed/model"
)

func (d *Domain) SignInUser(email, password string) error {
	user, err := d.GetUserByEmail(email)
	if err != nil {
		return err
	}

	if !d.util.CompareHashedPasswords(password, user.Password) {
		return apperrors.ErrPasswordMismatched
	}

	return nil
}

func (d *Domain) ChangePassword(email, password string) error {
	query := `UPDATE "user" SET password = $1 WHERE email = $2`

	_, err := d.pg.Exec(query, password, email)
	if err != nil {
		return apperrors.ErrPasswordUpdateFailed
	}

	return nil
}

func (d *Domain) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	query := `SELECT email, password, "loginIP", "isPasswordChanged" FROM "user" WHERE email=$1`
	err := d.pg.QueryRow(query, email).Scan(&user.Email, &user.Password, &user.LoginIP, &user.IsPasswordChanged)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, apperrors.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}
