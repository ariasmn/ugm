package tui

import (

	"github.com/charmbracelet/lipgloss"
)

func (bu BubbleUser) View() string {
	return docStyle.Render(bu.list.View())
}
