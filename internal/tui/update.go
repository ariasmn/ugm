package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (bu BubbleUser) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return bu, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		bu.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	bu.list, cmd = bu.list.Update(msg)
	return bu, cmd
}
