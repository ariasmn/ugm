package userparser

import (
	"reflect"
	"testing"
)

func TestGetUsers(t *testing.T) {
	var want []User
	got := GetUsers()

	if reflect.TypeOf(got) != reflect.TypeOf(want) {
        t.Errorf("got %q, wanted %q", reflect.TypeOf(got), reflect.TypeOf(want))
    }
}
