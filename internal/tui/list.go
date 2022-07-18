package tui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	user, ok := listItem.(item)
	if !ok {
		return
	}

	line := user.Details.Username

	if index == m.Index() {
		line = listSelectedlistItemStyle.Render("> " + line)
	} else {
		line = listItemStyle.Render(line)
	}

	fmt.Fprint(w, line)
}
