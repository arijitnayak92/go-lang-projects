package util

import "regexp"

var EmailRegex = "^[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+$"

func (u *Util) IsEmail(email string) (bool, error) {
	return regexp.MatchString(EmailRegex, email)
}
