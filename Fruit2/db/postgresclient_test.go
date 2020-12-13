package db

import "testing"

func TestConnectToPostgres(t *testing.T) {

	db, err := ConnectToPostgres("RandomString")

	if err != nil {
		t.Error("Error is :", db)
	}
}
