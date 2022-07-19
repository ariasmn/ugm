package tui

import (
	"github.com/ariasmn/ugm/internal/tui/group"
	"github.com/ariasmn/ugm/internal/tui/user"
	tea "github.com/charmbracelet/bubbletea"
)

type state int

const (
	showUserView state = iota
	showGroupView
)

type model struct {
	state state
	bu user.BubbleUser
	bg group.BubbleGroup
	width, height int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "tab" {
			return updateByState(m)
		}
	}

	switch m.state {
	case showUserView:
		m.bu, cmd = m.bu.Update(msg)
		return m, cmd
	case showGroupView:
		m.bg, cmd = m.bg.Update(msg)
		return m, cmd
	default:
		return m, nil
	}
}

func (m model) View() string {
	switch m.state {
	case showGroupView:
		return m.bg.View()
	default:
		return m.bu.View()
	}
}

func InitialModel() model {
	return model{
		state: showUserView, 
		bu: user.InitialModel(),
		bg: group.InitialModel(),
	}
}

func updateByState(m model) (model, tea.Cmd) {
	var cmd tea.Cmd
	windowSizeMsg := tea.WindowSizeMsg {
		Width: m.width,
		Height: m.height,
	}

	if m.state == showUserView {
		m.state = showGroupView
		m.bg, cmd = m.bg.Update(windowSizeMsg)
	} else {
		m.state = showUserView
		m.bu, cmd = m.bu.Update(windowSizeMsg)
	}

	return m, cmd
}
