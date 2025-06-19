package chooser

import (
	"cmp"
	"slices"

	"github.com/charmbracelet/bubbles/list"
)

// ...
type item struct {
	value   string
	current bool
}

// ...
func newItems(values []string, current string) []list.Item {
	items := make([]list.Item, 0, len(values))

	for _, v := range values {
		items = append(items, item{
			value:   v,
			current: isCurrent(current, v),
		})
	}

	//nolint:forcetypeassert // no risk of other types being here
	slices.SortFunc(items, func(a, b list.Item) int {
		return cmp.Compare(a.(item).value, b.(item).value)
	})

	return items
}

// ...
func (i item) String() string {
	return i.value
}

// ...
func (i item) Value() string {
	v := i.String()

	if i.current {
		return v + " *"
	}

	return v
}

// ...
func (i item) FilterValue() string {
	return i.String()
}

func isCurrent(a, b string) bool {
	if a == b {
		return true
	}

	if (a == "" && b == "default") || (a == "default" && b == "") {
		return true
	}

	return false
}
