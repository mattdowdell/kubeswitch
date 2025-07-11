package chooser

import (
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// ...
type Chooser struct {
	list   list.Model
	choice string
}

// ...
func New(title string, items []string, current string) *Chooser {
	return &Chooser{
		list: newList(title, items, current),
	}
}

func (c *Chooser) Choice() (string, bool) {
	if c.choice == "" {
		return "", false
	}

	return strings.TrimSuffix(c.choice, " *"), true
}

// ...
func (c *Chooser) Init() tea.Cmd {
	return nil
}

// ...
func (c *Chooser) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		c.list.SetSize(msg.Width, msg.Height)
		return c, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			return c, tea.Quit

		case "enter":
			i, ok := c.list.SelectedItem().(item)
			if ok {
				c.choice = i.String()
			}

			return c, tea.Quit
		}
	}

	var cmd tea.Cmd
	c.list, cmd = c.list.Update(msg)
	return c, cmd
}

// ...
func (c *Chooser) View() string {
	return "\n" + c.list.View()
}
