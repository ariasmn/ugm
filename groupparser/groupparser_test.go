//go:build linux || freebsd || openbsd || netbsd
// +build linux freebsd openbsd netbsd

package groupparser

import (
	"os"
	"os/user"
	"reflect"
	"testing"
)

const mockGroup = `test:x:0:
mock:x:65537:
`

func TestGetGroups(t *testing.T) {
	tempDir, _ := os.CreateTemp("", "group")
	err := os.WriteFile(tempDir.Name(), []byte(mockGroup), 0644)
	if err != nil {
		t.Fatalf("Should not have failed: %s", err)
	}
	defer os.Remove(tempDir.Name())

	ParseGroups(tempDir.Name())
	got := GetGroups()
	want := getWanted()

	for groupIndex, _ := range got {
		if got[groupIndex].Details.Gid != want[groupIndex].Details.Gid {
			t.Errorf("Got GID %s, wanted %s", got[groupIndex].Details.Gid, want[groupIndex].Details.Gid)
		}
		if got[groupIndex].Details.Name != want[groupIndex].Details.Name {
			t.Errorf("Got name %s, wanted %s", got[groupIndex].Details.Name, want[groupIndex].Details.Name)
		}
		if !reflect.DeepEqual(got[groupIndex].Users, want[groupIndex].Users) {
			t.Errorf("Got users %v, wanted %v", got[groupIndex].Users, want[groupIndex].Users)
		}
	}
}

func getWanted() []Group {
	testGroup := Group{}
	testGroup.Details.Gid = "0"
	testGroup.Details.Name = "test"
	testGroup.Users = []*user.User{nil}

	mockGroup := Group{}
	mockGroup.Details.Gid = "65537"
	mockGroup.Details.Name = "mock"
	mockGroup.Users = []*user.User{nil}

	// Hard to test the users inside the group, since we use the Golang's lookup method
	// This method tries to find an user within the system, so is quite hard to mock and not worth the hassle
	return []Group{testGroup, mockGroup}
}
