package utils

import (
	"os"
	"time"

	"github.com/arijitnayak92/taskAfford/Fruit/apperrors"
	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

// TokenDetails ...
type TokenDetails struct {
	RefreshToken string `json:"refresh_token" bson:"refresh_token,omitempty"`
	RefreshUUID  string `json:"refresh_uuid" bson:"refresh_uuid,omitempty"`
	RtExpires    int64  `json:"rt_expires" bson:"rt_expires,omitempty"`
}

// CreateToken ...
func (u *Util) CreateToken(email string) (*TokenDetails, error) {
	td := &TokenDetails{}

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUUID = uuid.NewV4().String() + "++" + email
	var err error

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["user_id"] = email
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func (u *Util) VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, apperrors.ErrWrongJwtMethod
		}
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (u *Util) TokenValid(t string) error {
	token, err := u.VerifyToken(t)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok || !token.Valid {
		return apperrors.ErrTokenExpired
	}
	return nil
}
