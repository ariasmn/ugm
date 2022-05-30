package tui

import (

	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)
func (bu BubbleUser) View() string {
	return docStyle.Render(bu.list.View())
}
