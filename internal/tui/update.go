package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (bu BubbleUser) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q":
            return bu, tea.Quit

        case "up", "k":
            if bu.cursor > 0 {
                bu.cursor--
            }

        case "down", "j":
            if bu.cursor < len(bu.users)-1 {
                bu.cursor++
            }
			
        case "enter", " ":
            _, ok := bu.selected[bu.cursor]
            if ok {
                delete(bu.selected, bu.cursor)
            } else {
                bu.selected[bu.cursor] = struct{}{}
            }
        }
    }

    return bu, nil
}
