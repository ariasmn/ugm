package user

import (
	"github.com/ariasmn/ugm/internal/tui/common"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

func (bu BubbleUser) Update(msg tea.Msg) (BubbleUser, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := common.ListStyle.GetFrameSize()
		bu.list.SetSize(msg.Width-h, msg.Height-v)
		bu.viewport = viewport.New(msg.Width, msg.Height)
		bu.viewport.SetContent(bu.detailView())
	}

	var cmd tea.Cmd
	bu.list, cmd = bu.list.Update(msg)
	
	return bu, cmd
}
