package domain

import (
	"testing"
)

func TestCreateUser(t *testing.T) {
	type testPerimeter struct {
		inputs   *User
		want     uint64
		hasError bool
	}
	testedData := []testPerimeter{
		{&User{001, "arijit", "nayak"}, 001, false},
		{&User{002, "gornf", "fort"}, 002, false},
		{&User{003, "arijit", "nayak"}, 003, true},
	}
	for _, user := range testedData {
		result, err := UserMethodMux.CreateUser(user.inputs)
		if user.hasError {
			if err == nil {
				t.Errorf("Creating User with data %v : FAILED, expected an error got an value %v", user.inputs, result)
			} else {
				t.Errorf("Creating User with data %v : PASSED, expected an error got an error %v", user.inputs, err)
			}
		} else {
			if result != user.want {
				t.Errorf("Creating User with data %v : FAILED, expected value %v got an value %v", user.inputs, user.want, result)
			} else {
				t.Errorf("Creating User with data %v : PASSED, expected value %v got an value %v", user.inputs, user.want, result)
			}
		}
	}
}
