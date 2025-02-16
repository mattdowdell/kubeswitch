package chooser

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

const (
	colorBlack = "0"
	colorCyan  = "6"
	colorWhite = "7"
)

//nolint:mnd // layout sizes
var (
	titleStyle      = lipgloss.NewStyle().MarginLeft(2)
	paginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle       = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)

	selectedColor = lipgloss.Color(colorCyan)
	textColor     = lipgloss.AdaptiveColor{
		Light: colorBlack,
		Dark:  colorWhite,
	}

	activeDot   = list.DefaultStyles().ActivePaginationDot.Foreground(selectedColor).String()
	inactiveDot = list.DefaultStyles().InactivePaginationDot.Foreground(textColor).String()
)

func newList(title string, items []string) list.Model {
	//nolint:mnd // layout sizes
	l := list.New(newItems(items), newDelegate(), 20 /*width*/, 30 /*height*/)

	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)

	l.Title = title
	l.InfiniteScrolling = true

	l.FilterInput.Cursor.Style = l.FilterInput.Cursor.Style.Foreground(textColor)
	l.FilterInput.PromptStyle = l.FilterInput.PromptStyle.Foreground(textColor)

	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	l.Paginator.ActiveDot = activeDot
	l.Paginator.InactiveDot = inactiveDot

	l.Help.Styles.ShortKey = l.Help.Styles.ShortKey.Foreground(textColor)
	l.Help.Styles.ShortDesc = l.Help.Styles.ShortDesc.Foreground(textColor)
	l.Help.Styles.ShortSeparator = l.Help.Styles.ShortSeparator.Foreground(textColor)
	l.Help.Styles.Ellipsis = l.Help.Styles.Ellipsis.Foreground(textColor)
	l.Help.Styles.FullKey = l.Help.Styles.FullKey.Foreground(textColor)
	l.Help.Styles.FullDesc = l.Help.Styles.FullDesc.Foreground(textColor)
	l.Help.Styles.FullSeparator = l.Help.Styles.FullSeparator.Foreground(textColor)

	return l
}
