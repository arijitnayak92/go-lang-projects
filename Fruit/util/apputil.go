package util

type AppUtil interface {
	HashPassword(password string) (string, error)
	CompareHashedPasswords(password, hashedPassword string) bool
	IsEmail(email string) (bool, error)
}

type Util struct {
}

func NewUtil() *Util {
	return &Util{}
}
