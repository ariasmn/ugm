//go:build linux || freebsd || openbsd || netbsd
// +build linux freebsd openbsd netbsd

package groupparser

import (
	"fmt"
	"os"
	"os/user"
	"reflect"
	"testing"
)

func TestGetGroups(t *testing.T) {
	currentUser, err := user.Current()
	if err != nil {
		t.Fatalf("Could not look up current user: %v", err)
	}

	tests := []struct {
		name  string
		group string
		want  []Group
	}{
		{
			name:  "empty member list",
			group: "empty:x:1000:\n",
			want: []Group{
				{
					Details: user.Group{Gid: "1000", Name: "empty"},
					Users:   []user.User{},
				},
			},
		},
		{
			name:  "resolves member",
			group: fmt.Sprintf("resolved:x:1001:%s\n", currentUser.Username),
			want: []Group{
				{
					Details: user.Group{Gid: "1001", Name: "resolved"},
					Users:   []user.User{*currentUser},
				},
			},
		},
		{
			name:  "keeps unresolved member",
			group: "unresolved:x:1002:ugm-test-unresolved-user\n",
			want: []Group{
				{
					Details: user.Group{Gid: "1002", Name: "unresolved"},
					Users:   []user.User{{Username: "ugm-test-unresolved-user"}},
				},
			},
		},
		{
			name:  "keeps mixed resolved and unresolved members",
			group: fmt.Sprintf("mixed:x:1003:%s,ugm-test-unresolved-user\n", currentUser.Username),
			want: []Group{
				{
					Details: user.Group{Gid: "1003", Name: "mixed"},
					Users: []user.User{
						*currentUser,
						{Username: "ugm-test-unresolved-user"},
					},
				},
			},
		},
		{
			name:  "skips empty member entries",
			group: fmt.Sprintf("empty-entries:x:1004:,%s,\n", currentUser.Username),
			want: []Group{
				{
					Details: user.Group{Gid: "1004", Name: "empty-entries"},
					Users:   []user.User{*currentUser},
				},
			},
		},
		{
			name: "parses multiple groups",
			group: "first:x:1005:\n" +
				"second:x:1006:ugm-test-unresolved-user\n",
			want: []Group{
				{
					Details: user.Group{Gid: "1005", Name: "first"},
					Users:   []user.User{},
				},
				{
					Details: user.Group{Gid: "1006", Name: "second"},
					Users:   []user.User{{Username: "ugm-test-unresolved-user"}},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempFile, err := os.CreateTemp("", "group")
			if err != nil {
				t.Fatalf("Could not create temp file: %v", err)
			}
			defer os.Remove(tempFile.Name())

			if _, err := tempFile.WriteString(tt.group); err != nil {
				t.Fatalf("Could not write temp file: %v", err)
			}
			if err := tempFile.Close(); err != nil {
				t.Fatalf("Could not close temp file: %v", err)
			}

			parseGroups(tempFile.Name())
			got := GetGroups()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGroups() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
