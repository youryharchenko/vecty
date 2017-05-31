package components

import (
	"github.com/youryharchenko/vecty"
	"github.com/youryharchenko/vecty/elem"
	"github.com/youryharchenko/vecty/event"
	"github.com/youryharchenko/vecty/examples/todomvc/actions"
	"github.com/youryharchenko/vecty/examples/todomvc/dispatcher"
	"github.com/youryharchenko/vecty/examples/todomvc/store"
	"github.com/youryharchenko/vecty/examples/todomvc/store/model"
	"github.com/youryharchenko/vecty/prop"
)

// FilterButton is a vecty.Component which allows the user to select a filter
// state.
type FilterButton struct {
	vecty.Core

	Label  string
	Filter model.FilterState
}

func (b *FilterButton) onClick(event *vecty.Event) {
	dispatcher.Dispatch(&actions.SetFilter{
		Filter: b.Filter,
	})
}

// Render implements the vecty.Component interface.
func (b *FilterButton) Render() *vecty.HTML {
	return elem.ListItem(
		elem.Anchor(
			vecty.If(store.Filter == b.Filter, prop.Class("selected")),
			prop.Href("#"),
			event.Click(b.onClick).PreventDefault(),

			vecty.Text(b.Label),
		),
	)
}
