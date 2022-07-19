package user

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

func (bu BubbleUser) Update(msg tea.Msg) (BubbleUser, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		bu.list.SetSize(msg.Width, msg.Height)
		bu.viewport = viewport.New(msg.Width, msg.Height)
		bu.viewport.SetContent(bu.detailView())
	}

	var cmd tea.Cmd
	bu.list, cmd = bu.list.Update(msg)
	
	return bu, cmd
}
