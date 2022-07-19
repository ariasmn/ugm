package group

import (
	"fmt"
	"os/user"
	"strings"

	"github.com/ariasmn/ugm/internal/tui/common"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
	"github.com/muesli/reflow/wordwrap"
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
	builder := &strings.Builder{}
	divider := common.DividerStyle.Render(strings.Repeat("-", bg.viewport.Width)) + "\n"
	detailsHeader := common.HeaderStyle.Render("Details")
	memberOfHeader := common.HeaderStyle.Render("Current members")

	if it := bg.list.SelectedItem(); it != nil {
		builder.WriteString(detailsHeader)
		builder.WriteString(renderGroupDetails(it.(item).Details))
		builder.WriteString(divider)
		builder.WriteString(memberOfHeader)
		builder.WriteString(fmt.Sprintf("\n\n%s", renderUserTable(it.(item).Users)))
	}
	details := wordwrap.String(builder.String(), bg.viewport.Width)

	return common.DetailStyle.Render(details)
}

func renderGroupDetails(group user.Group) string {
	gid := fmt.Sprintf("\n\nGID: %s\n", group.Gid)
	name := fmt.Sprintf("Name: %s\n", group.Name)

	return gid+name
}

func renderUserTable(users []*user.User) string {
	rows := []table.Row{}

	for _, user := range users {
		if user == nil {
			return "No users in this group"
		}
		rows = append(rows, table.NewRow(table.RowData{
			"Username"			: user.Username,
			"Fullname"			: user.Name,
			"UID"				: user.Uid,
			"GID"				: user.Gid,
			"Home directory"	: user.HomeDir,
		}))
	}

	userTable := table.New([]table.Column{
		table.NewColumn("Username", "Username", 16),
		table.NewColumn("Fullname", "Fullname", 20),
		table.NewColumn("UID", "UID", 10),
		table.NewColumn("GID", "GID", 10),
		table.NewColumn("Home directory", "Home directory", 25),
	}).WithRows(rows).
		BorderRounded().
		WithBaseStyle(common.TableMainStyle).
		HeaderStyle(common.TableHeaderStyle)

	return userTable.View()
}
