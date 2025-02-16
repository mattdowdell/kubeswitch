package chooser

import (
	"github.com/charmbracelet/bubbles/list"
)

// ...
type item string

// ...
func newItems(values []string) []list.Item {
	items := make([]list.Item, 0, len(values))

	for _, v := range values {
		items = append(items, item(v))
	}

	return items
}

// ...
func (i item) String() string {
	return string(i)
}

// ...
func (i item) FilterValue() string {
	return i.String()
}
