// +build linux freebsd openbsd netbsd

package userparser

import (
	"bufio"
	"errors"
	"log"
	"os"
	"os/user"
	"strings"
)

type User struct {
	Details user.User
	Groups []*user.Group
}

var parsedUsers []User

func GetUsers() (users []User) {
	if len(parsedUsers) == 0 {
		ParseUsers()
	}

	return parsedUsers
}

func ParseUsers() {
	parsedUsers = nil

	f, err := os.Open("/etc/passwd")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		user, err := parseLine(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		parsedUsers = append(parsedUsers, user)
	}
}

func parseLine(line string) (User, error) {
	fs := strings.Split(line, ":")

	if len(fs) != 7 {
		return User{}, errors.New("unexpected number of fields in /etc/passwd")
	}

	//Parse the GECOS field
	gecos := strings.Split(fs[4], ",")

	user := User{}
	user.Details.Uid = fs[2]
	user.Details.Gid = fs[3]
	user.Details.Username = fs[0]
	user.Details.Name = gecos[0]
	user.Details.HomeDir = fs[5]
	user.Groups = parseGroups(user.Details)

	return user, nil
}

func parseGroups(currentUser user.User) ([]*user.Group) {
	groupsIds, _:= currentUser.GroupIds()
	groups := []*user.Group{}

	for _, groupId := range groupsIds {
		foundGroup, _ := user.LookupGroupId(groupId)
		groups = append(groups, foundGroup)
	}

	return groups
}
