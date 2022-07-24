//go:build linux || freebsd || openbsd || netbsd
// +build linux freebsd openbsd netbsd

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
