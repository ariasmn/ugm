package group

import (
	"github.com/ariasmn/ugm/internal/tui/common"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (bg BubbleGroup) Update (msg tea.Msg) (BubbleGroup, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		horizontal, vertical := common.ListStyle.GetFrameSize()
		paginatorHeight := lipgloss.Height(bg.list.Paginator.View())

		bg.list.SetSize(msg.Width-horizontal, msg.Height-vertical-paginatorHeight)
		bg.viewport = viewport.New(msg.Width, msg.Height)
		bg.viewport.SetContent(bg.detailView())
	}

	var cmd tea.Cmd
	bg.list, cmd = bg.list.Update(msg)

	return bg, cmd
}
