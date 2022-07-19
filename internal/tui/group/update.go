package group

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

func (bg BubbleGroup) Update (msg tea.Msg) (BubbleGroup, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return bg, tea.Quit
		}
	case tea.WindowSizeMsg:
		bg.list.SetSize(msg.Width, msg.Height)
		bg.viewport = viewport.New(msg.Width, msg.Height)
		bg.viewport.SetContent(bg.detailView())
	}

	var cmd tea.Cmd
	bg.list, cmd = bg.list.Update(msg)

	return bg, cmd
}
