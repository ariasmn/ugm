//go:build linux || freebsd || openbsd || netbsd
// +build linux freebsd openbsd netbsd

package userparser

import (
	"os"
	"os/user"
	"reflect"
	"testing"
)

func TestGetUsers(t *testing.T) {
	rootGroup, err := user.LookupGroupId("0")
	if err != nil {
		t.Fatalf("Could not look up root group: %v", err)
	}

	tests := []struct {
		name   string
		passwd string
		want   []User
	}{
		{
			name:   "resolves primary group",
			passwd: "ugm-test-user:x:1000:0:ugm Test User:/home/ugm-test-user:/bin/sh\n",
			want: []User{
				{
					Details: user.User{
						Uid:      "1000",
						Gid:      "0",
						Username: "ugm-test-user",
						Name:     "ugm Test User",
						HomeDir:  "/home/ugm-test-user",
					},
					Groups: []user.Group{*rootGroup},
				},
			},
		},
		{
			name:   "keeps unresolved group",
			passwd: "ugm-test-user:x:1000:65537:ugm Test User:/home/ugm-test-user:/bin/sh\n",
			want: []User{
				{
					Details: user.User{
						Uid:      "1000",
						Gid:      "65537",
						Username: "ugm-test-user",
						Name:     "ugm Test User",
						HomeDir:  "/home/ugm-test-user",
					},
					Groups: []user.Group{{Gid: "65537"}},
				},
			},
		},
		{
			name:   "falls back to primary group when group IDs cannot be listed",
			passwd: ":x:1000:0:No Username:/home/no-username:/bin/sh\n",
			want: []User{
				{
					Details: user.User{
						Uid:      "1000",
						Gid:      "0",
						Username: "",
						Name:     "No Username",
						HomeDir:  "/home/no-username",
					},
					Groups: []user.Group{*rootGroup},
				},
			},
		},
		{
			name:   "skips empty group ID",
			passwd: "no-groups:x:1000::No Groups:/home/no-groups:/bin/sh\n",
			want: []User{
				{
					Details: user.User{
						Uid:      "1000",
						Gid:      "",
						Username: "no-groups",
						Name:     "No Groups",
						HomeDir:  "/home/no-groups",
					},
					Groups: []user.Group{},
				},
			},
		},
		{
			name: "parses multiple users",
			passwd: "resolved:x:1000:0:Resolved User:/home/resolved:/bin/sh\n" +
				"unresolved:x:1001:65537:Unresolved User:/home/unresolved:/bin/sh\n",
			want: []User{
				{
					Details: user.User{
						Uid:      "1000",
						Gid:      "0",
						Username: "resolved",
						Name:     "Resolved User",
						HomeDir:  "/home/resolved",
					},
					Groups: []user.Group{*rootGroup},
				},
				{
					Details: user.User{
						Uid:      "1001",
						Gid:      "65537",
						Username: "unresolved",
						Name:     "Unresolved User",
						HomeDir:  "/home/unresolved",
					},
					Groups: []user.Group{{Gid: "65537"}},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempFile, err := os.CreateTemp("", "passwd")
			if err != nil {
				t.Fatalf("Could not create temp file: %v", err)
			}
			defer os.Remove(tempFile.Name())

			if _, err := tempFile.WriteString(tt.passwd); err != nil {
				t.Fatalf("Could not write temp file: %v", err)
			}
			if err := tempFile.Close(); err != nil {
				t.Fatalf("Could not close temp file: %v", err)
			}

			parseUsers(tempFile.Name())
			got := GetUsers()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsers() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
