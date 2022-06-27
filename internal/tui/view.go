package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

func (bu BubbleUser) listView() string {
	return listStyle.Render(bu.list.View())
}

func (bu BubbleUser) detailView() string {
	builder := &strings.Builder{}
	divider := dividerStyle.Render(strings.Repeat("-", bu.viewport.Width)) + "\n"

	if it := bu.list.SelectedItem(); it != nil {
		username := fmt.Sprintf("Username: %s\n", it.(item).Username)
		fullname := fmt.Sprintf("Fullname: %s\n", it.(item).Name)
		identificators := fmt.Sprintf("UID: %s, GID: %s\n", it.(item).Uid, it.(item).Gid)
		homeDirectory := fmt.Sprintf("Home directory: %s\n", it.(item).HomeDir)

		builder.WriteString(username)
		builder.WriteString(fullname)
		builder.WriteString(identificators)
		builder.WriteString(homeDirectory)
		builder.WriteString(divider)
	}

	return wordwrap.String(builder.String(), bu.viewport.Width)
}

func (bu BubbleUser) View() string {
	bu.viewport.SetContent(bu.detailView())

	return lipgloss.JoinHorizontal(
		lipgloss.Top, bu.listView(), bu.viewport.View())
}
