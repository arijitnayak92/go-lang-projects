package domain

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/arijitnayak92/taskAfford/RESTMUXJWT/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v7"
	"github.com/twinj/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var (
	client        *redis.Client
	UserMethodMux userInterface
)

var userCollection = db().Database("goAPI").Collection("loginDB")

func init() {
	UserMethodMux = &usersStruct{}
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
	CreateUser(userDATA *User) (interface{}, *utils.APIError)
	Login(userDATA *User) (map[string]string, *utils.APIError)
	CreateToken(userid uint64) (*TokenDetails, *utils.APIError)
	CreateAuth(userid uint64, td *TokenDetails) *utils.APIError
	ExtractToken(r *http.Request) string
	VerifyToken(r *http.Request) (*jwt.Token, *utils.APIError)
	TokenValid(r *http.Request) *utils.APIError
	FetchAuth(authD *AccessDetails) (uint64, *utils.APIError)
	FetchUserID(req *http.Request) (uint64, *utils.APIError)
	ExtractTokenMetadata(*http.Request) (*AccessDetails, *utils.APIError)
	RefreshToken(req *http.Request) (map[string]string, *utils.APIError)
	DeleteAuth(givenUUID string) (int64, *utils.APIError)
	LogoutUser(req *http.Request) (int, *utils.APIError)
}

type usersStruct struct{}

func (c *usersStruct) CreateUser(userDATA *User) (interface{}, *utils.APIError) {
	var user *User
	userCollection.FindOne(context.TODO(), bson.M{"username": userDATA.Username}).Decode(&user)
	if user != nil {
		return nil, &utils.APIError{
			Message:    "User already present !",
			StatusCode: 422,
		}
	}

	bytes, errs := bcrypt.GenerateFromPassword([]byte(userDATA.Password), 14)
	if errs != nil {
		return nil, &utils.APIError{
			Message:    "Something went wrong !",
			StatusCode: 422,
		}
	}
	userDATA.Password = string(bytes)
	addedUser, err := userCollection.InsertOne(context.TODO(), userDATA)
	if err != nil {
		return nil, &utils.APIError{
			Message:    "Something went wrong !",
			StatusCode: 422,
		}
	}
	fmt.Println(addedUser.InsertedID)
	return addedUser.InsertedID, nil
}

func (c *usersStruct) Login(userDATA *User) (map[string]string, *utils.APIError) {
	var user *User
	if err := userCollection.FindOne(context.TODO(), bson.M{"username": userDATA.Username}).Decode(&user); err != nil {
		return nil, &utils.APIError{
			Message:    "User Not Found !",
			StatusCode: 401,
		}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDATA.Password)); err != nil {
		return nil, &utils.APIError{
			Message:    "Wrong Password !",
			StatusCode: 401,
		}
	}
	ts, err := UserMethodMux.CreateToken(user.Id)
	if err != nil {
		return nil, &utils.APIError{
			Message:    "Something went wrong !",
			StatusCode: 422,
		}
	}
	saveErr := UserMethodMux.CreateAuth(user.Id, ts)
	if saveErr != nil {
		return nil, &utils.APIError{
			Message:    "Something went wrong !",
			StatusCode: 422,
		}
	}
	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}
	return tokens, nil
}

func (t *usersStruct) CreateToken(userid uint64) (*TokenDetails, *utils.APIError) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = td.AccessUuid + "++" + strconv.Itoa(int(userid))
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
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
			StatusCode: 422,
		}
	}
	//Creating Refresh Token
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf")
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, &utils.APIError{
			Message:    "Something went wrong !",
			StatusCode: 422,
		}
	}
	return td, nil
}

func (t *usersStruct) CreateAuth(userid uint64, td *TokenDetails) *utils.APIError {
	at := time.Unix(td.AtExpires, 0)
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

func (t *usersStruct) ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	return bearToken
}

// this function will receive the parsed token and should return the key for validating.
func (t *usersStruct) VerifyToken(r *http.Request) (*jwt.Token, *utils.APIError) {
	tokenString := UserMethodMux.ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, &utils.APIError{
			Message:    "Something went wrong !",
			StatusCode: 422,
		}
	}
	return token, nil
}

func (t *usersStruct) TokenValid(r *http.Request) *utils.APIError {
	token, err := UserMethodMux.VerifyToken(r)
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

func (t *usersStruct) FetchUserID(req *http.Request) (uint64, *utils.APIError) {
	//Extract the access token metadata
	metadata, err := UserMethodMux.ExtractTokenMetadata(req)
	if err != nil {
		return 0, &utils.APIError{
			Message:    "UnAuthorized  !",
			StatusCode: 401,
		}
	}
	userid, err := UserMethodMux.FetchAuth(metadata)
	if err != nil {
		return 0, &utils.APIError{
			Message:    "UnAuthorized  !",
			StatusCode: 401,
		}
	}
	return userid, nil
}

func (t *usersStruct) ExtractTokenMetadata(r *http.Request) (*AccessDetails, *utils.APIError) {
	token, err := UserMethodMux.VerifyToken(r)
	if err != nil {
		return nil, &utils.APIError{
			Message:    "Invalid Token  !",
			StatusCode: 401,
		}
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, &utils.APIError{
				Message:    "UnAuthorized  !",
				StatusCode: 401,
			}
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, &utils.APIError{
				Message:    "Something went wrong  !",
				StatusCode: 422,
			}
		}
		return &AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
		}, nil
	}
	return nil, &utils.APIError{
		Message:    "Something went wrong  !",
		StatusCode: 422,
	}
}

func (t *usersStruct) FetchAuth(authD *AccessDetails) (uint64, *utils.APIError) {
	userid, err := client.Get(authD.AccessUuid).Result()
	if err != nil {
		return 0, &utils.APIError{
			Message:    "Unauthorized Action !",
			StatusCode: 403,
		}
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	if authD.UserId != userID {
		return 0, &utils.APIError{
			Message:    "Unauthorized Action !",
			StatusCode: 403,
		}
	}
	return userID, nil
}

func (t *usersStruct) RefreshToken(req *http.Request) (map[string]string, *utils.APIError) {
	mapToken := map[string]string{}
	if err := json.NewDecoder(req.Body).Decode(&mapToken); err != nil {
		return nil, &utils.APIError{
			Message:    "JSON converting error !",
			StatusCode: 422,
		}
	}
	refreshToken := mapToken["refresh_token"]

	//verify the token
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf")
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})

	if err != nil {
		return nil, &utils.APIError{
			Message:    "Refresh token expired !",
			StatusCode: 401,
		}
	}
	//is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, &utils.APIError{
			Message:    "Unauthorized !",
			StatusCode: 401,
		}
	}
	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string)
		if !ok {
			return nil, &utils.APIError{
				Message:    "Unable to process data !",
				StatusCode: 422,
			}
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, &utils.APIError{
				Message:    "Some error occoured !",
				StatusCode: 422,
			}
		}
		//Delete the previous Refresh Token
		deleted, delErr := UserMethodMux.DeleteAuth(refreshUuid)
		if delErr != nil || deleted == 0 {
			return nil, &utils.APIError{
				Message:    "Unauthorized !",
				StatusCode: 401,
			}
		}
		//Create new pairs of refresh and access tokens
		ts, createErr := UserMethodMux.CreateToken(userId)
		if createErr != nil {
			return nil, &utils.APIError{
				Message:    "Denied !",
				StatusCode: 403,
			}
		}
		//save the tokens metadata to redis
		saveErr := UserMethodMux.CreateAuth(userId, ts)
		if saveErr != nil {
			return nil, &utils.APIError{
				Message:    "Denied !",
				StatusCode: 403,
			}
		}
		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		return tokens, nil
	} else {
		return nil, &utils.APIError{
			Message:    "Unauthorized !",
			StatusCode: 401,
		}
	}
}

func (t *usersStruct) DeleteAuth(givenUUID string) (int64, *utils.APIError) {
	deleted, err := client.Del(givenUUID).Result()
	if err != nil {
		return 0, &utils.APIError{
			Message:    "Error Occoured !",
			StatusCode: 422,
		}
	}
	return deleted, nil
}

func DeleteTokens(authD *AccessDetails) error {
	//get the refresh uuid
	refreshUuid := fmt.Sprintf("%s++%d", authD.AccessUuid, authD.UserId)
	fmt.Println("printing uuid")

	//delete access token
	deletedAt, err := client.Del(authD.AccessUuid).Result()
	if err != nil {
		return err
	}
	//delete refresh token
	deletedRt, err := client.Del(refreshUuid).Result()
	if err != nil {
		return err
	}

	if deletedAt != 1 || deletedRt != 1 {
		return errors.New("something went wrong")
	}
	return nil
}

func (t *usersStruct) LogoutUser(req *http.Request) (int, *utils.APIError) {
	metadata, err := UserMethodMux.ExtractTokenMetadata(req)
	if err != nil {
		return 0, &utils.APIError{
			Message:    "Unauthorized !",
			StatusCode: 401,
		}
	}
	fmt.Println(metadata)
	delErr := DeleteTokens(metadata)
	if delErr != nil {
		return 0, &utils.APIError{
			Message:    "Unauthorized !",
			StatusCode: 401,
		}
	}
	return 1, nil
}
