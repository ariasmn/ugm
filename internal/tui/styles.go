package tui

import "github.com/charmbracelet/lipgloss"

var (
	listStyle = lipgloss.NewStyle().
			Width(28).
			PaddingRight(3).
			MarginRight(3).
			Border(lipgloss.RoundedBorder(), false, true, false, false)
	dividerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#9B9B9B", Dark: "#5C5C5C"})
)
