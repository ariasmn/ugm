//go:build linux || freebsd || openbsd || netbsd
// +build linux freebsd openbsd netbsd

package userparser

import (
	"os"
	"os/user"
	"reflect"
	"testing"
)

const mockPwd = `test:x:0:0:test:/test:/bin/test
mock:x:1:65537:mock:/mock:/sbin/mock
`

func TestGetUsers(t *testing.T) {
	tempDir, _ := os.CreateTemp("", "passwd")
	err := os.WriteFile(tempDir.Name(), []byte(mockPwd), 0644)
	if err != nil {
		t.Fatalf("Should not have failed: %s", err)
	}
	defer os.Remove(tempDir.Name())

	ParseUsers(tempDir.Name()) // We parse the temp file before getting the users
	got := GetUsers()
	want := getWanted()

	for userIndex, _ := range got {
		if got[userIndex].Details.Uid != want[userIndex].Details.Uid {
			t.Errorf("Got UID %s, wanted %s", got[userIndex].Details.Uid, want[userIndex].Details.Uid)
		}
		if got[userIndex].Details.Gid != want[userIndex].Details.Gid {
			t.Errorf("Got GID %s, wanted %s", got[userIndex].Details.Gid, want[userIndex].Details.Gid)
		}
		if got[userIndex].Details.Username != want[userIndex].Details.Username {
			t.Errorf("Got username %s, wanted %s", got[userIndex].Details.Username, want[userIndex].Details.Username)
		}
		if got[userIndex].Details.Name != want[userIndex].Details.Name {
			t.Errorf("Got name %s, wanted %s", got[userIndex].Details.Name, want[userIndex].Details.Name)
		}
		if got[userIndex].Details.HomeDir != want[userIndex].Details.HomeDir {
			t.Errorf("Got homedir %s, wanted %s", got[userIndex].Details.HomeDir, want[userIndex].Details.HomeDir)
		}
		if !reflect.DeepEqual(got[userIndex].Groups, want[userIndex].Groups) {
			t.Errorf("Got group %v, wanted %v", got[userIndex].Groups, want[userIndex].Groups)
		}
	}
}

func getWanted() []User {
	userTest := User{}
	rootGroup, _ := user.LookupGroupId("0")

	userTest.Details.Uid = "0"
	userTest.Details.Gid = "0"
	userTest.Details.Username = "test"
	userTest.Details.Name = "test"
	userTest.Details.HomeDir = "/test"
	userTest.Groups = append(userTest.Groups, rootGroup)

	userMock := User{}
	userMock.Details.Uid = "1"
	userMock.Details.Gid = "65537"
	userMock.Details.Username = "mock"
	userMock.Details.Name = "mock"
	userMock.Details.HomeDir = "/mock"
	userMock.Groups = []*user.Group{nil}

	return []User{userTest, userMock}
}
