package tui

import (
	"github.com/ariasmn/ugm/userparser"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
)

type item userparser.User

func (i item) FilterValue() string { return i.Details.Username }

type BubbleUser struct {
	list list.Model
	viewport viewport.Model
}

func InitialModel() BubbleUser {
	items := userToItem(userparser.GetUsers())
	l := list.New(items, itemDelegate{}, 0, 0)
	l.Title = "Users"
	l.SetShowHelp(false)

	return BubbleUser{list: l}
}

func userToItem(users []userparser.User) []list.Item {
	items := make([]list.Item, len(users))
	for i,v := range users {
		items[i] = item(v)
	}

	return items
}
