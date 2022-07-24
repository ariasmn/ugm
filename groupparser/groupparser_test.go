//go:build linux || freebsd || openbsd || netbsd
// +build linux freebsd openbsd netbsd

package groupparser

import (
	"reflect"
	"testing"
)

func TestGetGroups(t *testing.T) {
	var want []Group
	got := GetGroups()

	if reflect.TypeOf(got) != reflect.TypeOf(want) {
		t.Errorf("got %q, wanted %q", reflect.TypeOf(got), reflect.TypeOf(want))
	}
}
