package tui

import "github.com/charmbracelet/lipgloss"

var (
	listStyle = lipgloss.NewStyle().
		Width(35).
		PaddingRight(3).
		MarginRight(3).
		Border(lipgloss.RoundedBorder(), false, true, false, false)
	listColorStyle = lipgloss.NewStyle().
		Background(lipgloss.Color("#3d719c"))
	detailStyle = lipgloss.NewStyle().
		PaddingTop(2)
	dividerStyle = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#9B9B9B", Dark: "#5C5C5C"}).
		PaddingTop(1).
		PaddingBottom(1)
	tableMainStyle = lipgloss.NewStyle().
		Align(lipgloss.Center)
	tableHeaderStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#569cd6")).
		Bold(true)
)
