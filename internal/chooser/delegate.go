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

func (*delegate) Height() int {
	return 1
}

func (*delegate) Spacing() int {
	return 0
}

func (*delegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd {
	return nil
}

// ...
//
//nolint:gocritic // (hugeParam) required interface defined elsewhere
func (*delegate) Render(writer io.Writer, model list.Model, index int, item list.Item) {
	i, ok := item.(fmt.Stringer)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i.String())

	fn := itemStyle.Render
	if index == model.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(writer, fn(str))
}
