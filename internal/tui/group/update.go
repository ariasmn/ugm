package group

import (
	"github.com/ariasmn/ugm/internal/tui/common"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

func (bg BubbleGroup) Update (msg tea.Msg) (BubbleGroup, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := common.ListStyle.GetFrameSize()
		bg.list.SetSize(msg.Width-h, msg.Height-v)
		bg.viewport = viewport.New(msg.Width, msg.Height)
		bg.viewport.SetContent(bg.detailView())
	}

	var cmd tea.Cmd
	bg.list, cmd = bg.list.Update(msg)

	return bg, cmd
}
