package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	//Initialization

	//Execuation
	user, err := UserMethods.GetUser(0)
	//Validation
	assert.Nil(t, user, "Don't want a user with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, 404, err.StatusCode)
	assert.EqualValues(t, "User Not Found !", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	//Initialization

	//Execuation
	user, err := UserMethods.GetUser(123)
	//Validation
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Arijit", user.FirstName)
	assert.EqualValues(t, "Nayak", user.LastName)
	assert.EqualValues(t, "arijitnayak92@gmail.com", user.Email)
}
