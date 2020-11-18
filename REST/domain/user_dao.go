package domain

import (
	"github.com/arijitnayak92/taskAfford/REST/utils"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"github.com/dgrijalva/jwt-go"
"github.com/go-redis/redis/v7"
	"github.com/twinj/uuid"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Arijit", LastName: "Nayak", Email: "arijitnayak92@gmail.com"},
	}
	client *redis.Client
	UserMethods userInterface
) 

func init() {
	UserMethods = &usersStruct{}
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	client = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

type userInterface interface {
	Login(userID int64) (*User, *utils.APIError)
	CreateToken(userid int64) (*TokenDetails, *utils.APIError)
	CreateAuth(userid int64, td *TokenDetails) *utils.APIError
	ExtractToken(r *http.Request) string
	VerifyToken(r *http.Request) (*jwt.Token, *utils.APIError)
	TokenValid(r *http.Request) *utils.APIError
	FetchAuth(authD *AccessDetails) (uint64,*utils.APIError)
}

type usersStruct struct{}

func (c *usersStruct) Login(userID int64) (*User, *utils.APIError) {
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.APIError{
		Message:    "User Not Found !",
		StatusCode: 404,
	}
}

func (t *usersStruct)CreateToken(userid int64) (*TokenDetails,*utils.APIError) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = td.AccessUuid + "++" + strconv.Itoa(int(userid))
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, &utils.APIError{
			Message:    "Something went wrong !",
			StatusCode: 406,
		}
	}
	//Creating Refresh Token
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf") //this should be in an env file
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, &utils.APIError{
			Message:    "Something went wrong !",
			StatusCode: 406,
		}
	}
	return td, nil
}

func (t *usersStruct)CreateAuth(userid int64, td *TokenDetails) *utils.APIError {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := client.Set(td.AccessUuid, strconv.Itoa(int(userid)), at.Sub(now)).Err()
	if errAccess != nil {
		return &utils.APIError{
			Message:    "Token Not Set !",
			StatusCode: 422,
		}
	}
	errRefresh := client.Set(td.RefreshUuid, strconv.Itoa(int(userid)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return &utils.APIError{
			Message:    "Token Not Set !",
			StatusCode: 422,
		}
	}
	return nil
}

func (t *usersStruct)ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// Parse, validate, and return a token.
// keyFunc will receive the parsed token and should return the key for validating.
func (t *usersStruct)VerifyToken(r *http.Request) (*jwt.Token, *utils.APIError) {
	tokenString := UserMethods.ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil,&utils.APIError{
			Message:    "Something went wrong !",
			StatusCode: 422,
		}
	}
	return token, nil
}

func (t *usersStruct)TokenValid(r *http.Request) *utils.APIError {
	token, err := UserMethods.VerifyToken(r)
	if err != nil {
		return &utils.APIError{
			Message:    "Invalid Token !",
			StatusCode: 401,
		}
	}
	if _, ok := token.Claims.(jwt.Claims); !ok || !token.Valid {
		return &utils.APIError{
			Message:    "Token Expired !",
			StatusCode: 401,
		}
	}
	return nil
}

func (t *usersStruct)ExtractTokenMetadata(r *http.Request) (*AccessDetails,*utils.APIError) {
	token, err := UserMethods.VerifyToken(r)
	if err != nil {
		return nil,&utils.APIError{
			Message:    "Invalid Token !",
			StatusCode: 401,
		}
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, &utils.APIError{
				Message:    "Token Expired !",
				StatusCode: 401,
			}
		}
		userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, &utils.APIError{
				Message:    "Something went wrong !",
				StatusCode: 422,
			}
		}
		return &AccessDetails{
			AccessUuid: accessUUID,
			UserId:     userID,
		}, nil
	}
	return nil, &utils.APIError{
		Message:    "Something went wrong !",
		StatusCode: 422,
	}
}

func (t *usersStruct)FetchAuth(authD *AccessDetails) (uint64,*utils.APIError) {
	userid, err := client.Get(authD.AccessUuid).Result()
	if err != nil {
		return 0, &utils.APIError{
			Message:    "Unauthorized Action !",
			StatusCode: 403,
		}
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	if authD.UserId != userID {
		return 0,&utils.APIError{
			Message:    "Unauthorized Action !",
			StatusCode: 403,
		}
	}
	return userID, nil
}
