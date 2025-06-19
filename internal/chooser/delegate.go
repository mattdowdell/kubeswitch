package chooser

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

//nolint:mnd // layout sizes
var (
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(selectedColor)
)

type delegate struct{}

func newDelegate() list.ItemDelegate {
	return &delegate{}
}

// Height is the height of the list item.
func (*delegate) Height() int {
	return 1
}

// Spacing is the size of the horizontal gap between list items in cells.
func (*delegate) Spacing() int {
	return 0
}

// Update is the update loop for items.
//
// All messages in the list's update loop will pass through here except when the user is setting a
// filter. Use this method to perform item-level updates appropriate to this delegate.
func (*delegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd {
	return nil
}

// Render renders the item's view.
//
//nolint:gocritic // (hugeParam) required interface defined elsewhere
func (*delegate) Render(writer io.Writer, model list.Model, index int, item list.Item) {
	i, ok := item.(interface{ Value() string })
	if !ok {
		return
	}

	str := i.Value()

	fn := itemStyle.Render
	if index == model.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(writer, fn(str))
}
