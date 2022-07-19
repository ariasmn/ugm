package group

import (
	"github.com/ariasmn/ugm/groupparser"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
)

type item groupparser.Group

func (i item) FilterValue() string { return i.Details.Name }

type BubbleGroup struct {
	list list.Model
	viewport viewport.Model
}

func InitialModel() BubbleGroup {
	items := groupToItem(groupparser.GetGroups())
	l := list.New(items, itemDelegate{}, 0, 0)
	l.Title = "Groups"
	l.SetShowHelp(false)

	return BubbleGroup{list: l}
}

func groupToItem(groups []groupparser.Group) []list.Item {
	items := make([]list.Item, len(groups))
	for i,v := range groups {
		items[i] = item(v)
	}

	return items

}