package user

import (
	"fmt"
	"os/user"
	"strings"

	"github.com/ariasmn/ugm/internal/tui/common"
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
	bu.list.Styles.Title = common.ListColorStyle
	bu.list.Styles.FilterPrompt.Foreground(common.ListColorStyle.GetBackground())
	bu.list.Styles.FilterCursor.Foreground(common.ListColorStyle.GetBackground())

	return common.ListStyle.Render(bu.list.View())
}

func (bu BubbleUser) detailView() string {
	builder := &strings.Builder{}
	divider := common.DividerStyle.Render(strings.Repeat("-", bu.viewport.Width)) + "\n"
	detailsHeader := common.HeaderStyle.Render("Details")
	memberOfHeader := common.HeaderStyle.Render("Member of")

	if it := bu.list.SelectedItem(); it != nil {
		builder.WriteString(detailsHeader)
		builder.WriteString(renderUserDetails(it.(item).Details))
		builder.WriteString(divider)
		builder.WriteString(memberOfHeader)
		builder.WriteString(fmt.Sprintf("\n\n%s", renderGroupTable(it.(item).Groups)))
	}
	details := wordwrap.String(builder.String(), bu.viewport.Width)

	return common.DetailStyle.Render(details)
}

func renderUserDetails(user user.User) string {
	username := fmt.Sprintf("\n\nUsername: %s\n", user.Username)
	fullname := fmt.Sprintf("Fullname: %s\n", user.Name)
	identificators := fmt.Sprintf("UID: %s\nGID: %s\n", user.Uid, user.Gid)
	homeDirectory := fmt.Sprintf("Home directory: %s\n", user.HomeDir)

	return username + fullname + identificators + homeDirectory
}

func renderGroupTable(groups []*user.Group) string {
	rows := []table.Row{}

	for _, group := range groups {
		rows = append(rows, table.NewRow(table.RowData{
			"GID":  group.Gid,
			"Name": group.Name,
		}))
	}

	groupsTable := table.New([]table.Column{
		table.NewColumn("GID", "GID", 10),
		table.NewColumn("Name", "Name", 16),
	}).WithRows(rows).
		BorderRounded().
		WithBaseStyle(common.TableMainStyle).
		HeaderStyle(common.TableHeaderStyle)

	return groupsTable.View()
}
