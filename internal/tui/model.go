package tui

import (
	"github.com/ariasmn/ugm/userparser"
	"github.com/charmbracelet/bubbles/list"
)

type item userparser.User

func (i item) Title() string       { return i.Username }
func (i item) Description() string { return i.HomeDir }
func (i item) FilterValue() string { return i.Username }

type BubbleUser struct {
	list list.Model
}

func InitialModel() BubbleUser {
	items := userToItem(userparser.GetUsers())
	
	return BubbleUser{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
}

func userToItem(users []userparser.User) []list.Item {
	items := make([]list.Item, len(users))
	for i,v := range users {
		items[i] = item{Uid: v.Uid, Gid: v.Gid, Username: v.Username, HomeDir: v.HomeDir}
	}

	return items
}