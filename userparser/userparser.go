package userparser

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/user"
	"strings"
)

type User = user.User

func GetUsers() (users []User) {
	return ParseUsers()
}

func ParseUsers() (users []User) {
	f, err := os.Open("/etc/passwd")
	if err != nil {
		fmt.Errorf("%v", err)
	}

	defer f.Close()

	var usersSlice []User
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		user, err := parseLine(scanner.Text())
		if err != nil {
			fmt.Errorf("%v", err)
		}

		usersSlice = append(usersSlice, user)
	}

	return usersSlice
}

func parseLine(line string) (user.User, error) {
	fs := strings.Split(line, ":")

	if len(fs) != 7 {
		return user.User{}, errors.New("unexpected number of fields in /etc/passwd")
	}

	//Parse the GECOS field
	gecos := strings.Split(fs[4], ",")

	return user.User{
		Uid: fs[2],
		Gid: fs[3],
		Username: fs[0],
		Name: gecos[0],
		HomeDir: fs[5]}, nil
}
