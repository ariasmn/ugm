package tui

import (
	"fmt"
)

func (bu BubbleUser) View() string {
    ui := "System users\n\n"

    for i, user := range bu.users {

        cursor := " " // no cursor
        if bu.cursor == i {
            cursor = ">" // cursor!
        }

        ui += fmt.Sprintf("%s %s\n", cursor, user.Username)
    }

    ui += "\nPress q to quit.\n"

    return ui
}
