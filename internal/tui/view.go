package tui

func (bu BubbleUser) View() string {
	return docStyle.Render(bu.list.View())
}
