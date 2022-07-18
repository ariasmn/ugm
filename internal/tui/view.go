package tui

import (
	"fmt"
	"os/user"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
	"github.com/muesli/reflow/wordwrap"
)

func (bu BubbleUser) View() string {
	bu.viewport.SetContent(bu.detailView())

	return lipgloss.JoinHorizontal(
		lipgloss.Top, bu.listView(), bu.viewport.View())
}

func (bu BubbleUser) listView() string {
	bu.list.Styles.Title = listColorStyle

	return listStyle.Render(bu.list.View())
}

func (bu BubbleUser) detailView() string {
	builder := &strings.Builder{}

	divider := dividerStyle.Render(strings.Repeat("-", bu.viewport.Width)) + "\n"
	detailsHeader := headerStyle.Render("Details")
	memberOfHeader := headerStyle.Render("Member of")

	if it := bu.list.SelectedItem(); it != nil {
		username := fmt.Sprintf("\n\nUsername: %s\n", it.(item).Details.Username)
		fullname := fmt.Sprintf("Fullname: %s\n", it.(item).Details.Name)
		identificators := fmt.Sprintf("UID: %s\nGID: %s\n", it.(item).Details.Uid, it.(item).Details.Gid)
		homeDirectory := fmt.Sprintf("Home directory: %s\n", it.(item).Details.HomeDir)

		builder.WriteString(detailsHeader)
		builder.WriteString(username)
		builder.WriteString(fullname)
		builder.WriteString(identificators)
		builder.WriteString(homeDirectory)

		builder.WriteString(divider)

		builder.WriteString(memberOfHeader)
		builder.WriteString(fmt.Sprintf("\n\n%s", renderGroupTable(it.(item).Groups).View()))

	}
	details := wordwrap.String(builder.String(), bu.viewport.Width)

	return detailStyle.Render(details)
}

func renderGroupTable(groups []*user.Group) table.Model {
	rows := []table.Row{}

	for _, group := range groups {
		rows = append(rows, table.NewRow(table.RowData{
			"GID": group.Gid,
			"Name": group.Name,
		}))
	}

	groupsTable := table.New([]table.Column{
		table.NewColumn("GID", "GID", 10),
		table.NewColumn("Name", "Name", 16),
	}).WithRows(rows).
		BorderRounded().
		WithBaseStyle(tableMainStyle).
		HeaderStyle(tableHeaderStyle)

	return groupsTable
}
