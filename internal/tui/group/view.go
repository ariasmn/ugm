package group

import (
	"github.com/ariasmn/ugm/internal/tui/common"
	"github.com/charmbracelet/lipgloss"
)

func (bg BubbleGroup) View() string {
	bg.viewport.SetContent(bg.detailView())

	return lipgloss.JoinHorizontal(
		lipgloss.Top, bg.listView(), bg.viewport.View())
}

func (bg BubbleGroup) listView() string {
	bg.list.Styles.Title = common.ListColorStyle

	return common.ListStyle.Render(bg.list.View())
}

func (bg BubbleGroup) detailView() string {
	return "TODO"
}
