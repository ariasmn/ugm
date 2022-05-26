package tui

import (
	"github.com/ariasmn/ugm/userparser"
)

type BubbleUser struct {
	users []userparser.User
	cursor int
	selected map[int]struct{}
}

func InitialModel() BubbleUser {
	return BubbleUser{
		users: userparser.GetUsers(),
		selected: make(map[int]struct{}),
	}
}
