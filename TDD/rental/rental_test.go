package rental

import (
	"reflect"
	"testing"

	"github.com/taskAfford/TDD/receipt"
)

func TestUserExist(t *testing.T) {
	newRental := NewRental()
	want := false
	got := newRental.CheckUser("Arijit")
	if got != want {
		t.Errorf("Checkuser expection mismatched ! want %v got %v", want, got)
	}

	want = true
	newRental.AddUser("Arijits")
	got = newRental.CheckUser("Arijits")
	if got != want {
		t.Errorf("Checkuser expection mismatched ! want %v got %v", want, got)
	}

	t.Log("Testing with bulk of existing user")
	{
		newRental := NewRental()
		want := true
		users := []string{"Aryansh", "Bijendra", "Vivek"}
		for _, user := range users {
			newRental.AddUser(user)
			got := newRental.CheckUser(user)
			if got != want {
				t.Errorf("Checkuser expection mismatched ! want %v got %v", want, got)
			}
		}
	}
}

func TestRent(t *testing.T) {
	t.Log("Receipt information matches")
	{
		want := &receipt.Receipt{
			movies: []Movie{
				{
					Name: "Avengers EndGame",
					Date: "02-05-2020",
				},
				{
					Name: "Avengers EndGame 2",
					Date: "02-05-2020",
				},
			},
		}
		//newReceipt:= NewReceipt()
		newRental := NewRental()
		got := newReceipt.GenerateReceipt()

		if !reflect.DeepEqual(got, want) {
			t.Error(got, want)
		}
	}
}
