package db

import (
	"testing"
)

func TestMongoHealthCheck(t *testing.T) {

	_ = NewRepository(nil, nil)

	// err := newRepo.MongoHealthCheck()

	// t.Log(err)

}

func TestConnectToMongo(t *testing.T) {

	client, err := ConnectToMongo("RandomString")

	if err == nil {
		t.Error("Error is :", client)
	}
}
