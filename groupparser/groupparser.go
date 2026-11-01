//go:build linux || freebsd || openbsd || netbsd
// +build linux freebsd openbsd netbsd

package groupparser

import (
	"bufio"
	"errors"
	"log"
	"os"
	"os/user"
	"strings"
)

type Group struct {
	Details user.Group
	Users   []user.User
}

var parsedGroups []Group

func GetGroups() (groups []Group) {
	if len(parsedGroups) == 0 {
		parseGroups("/etc/group")
	}

	return parsedGroups
}

func parseGroups(path string) {
	parsedGroups = nil

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		group, err := parseLine(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		parsedGroups = append(parsedGroups, group)
	}
}

func parseLine(line string) (Group, error) {
	fs := strings.Split(line, ":")

	if len(fs) != 4 {
		return Group{}, errors.New("unexpected number of fields in /etc/group")
	}

	group := Group{}
	group.Details.Gid = fs[2]
	group.Details.Name = fs[0]
	group.Users = parseUsers(fs[3])

	return group, nil
}

func parseUsers(groupUsernames string) []user.User {
	usernames := strings.Split(groupUsernames, ",")
	users := make([]user.User, 0, len(usernames))

	for _, username := range usernames {
		if username == "" {
			continue
		}

		foundUser, err := user.Lookup(username)
		if err != nil {
			// Keep membership with username and empty details.
			users = append(users, user.User{Username: username})
			continue
		}

		users = append(users, *foundUser)
	}

	return users
}
