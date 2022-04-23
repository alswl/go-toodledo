package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"testing"
)

func TestSidebar(t *testing.T) {
	items := []list.Item{
		item{title: "Raspberry Pi’s"},
		item{title: "Nutella"},
		item{title: "Bitter melon"},
		item{title: "Nice socks"},
		item{title: "Eight hours of sleep"},
		item{title: "Cats"},
		item{title: "Plantasia, the album"},
		item{title: "Pour over coffee"},
		item{title: "VR"},
		item{title: "Noguchi Lamps"},
		item{title: "Linux"},
		item{title: "Business school"},
		item{title: "Pottery"},
		item{title: "Shampoo"},
		item{title: "Table tennis"},
		item{title: "Milk crates"},
		item{title: "Afternoon tea"},
		item{title: "Stickers"},
		item{title: "20° Weather"},
		item{title: "Warm light"},
		item{title: "The vernal equinox"},
		item{title: "Gaffer’s tape"},
		item{title: "Terrycloth"},
	}

	m := SidebarPane{
		isCollapsed: false,
		tabs:        []string{"foo", "bar", "baz"},
		currentTab:  "bar",
		items:       []string{},
		list:        list.New(items, list.NewDefaultDelegate(), 0, 0),
		currentItem: "",
	}
	view := m.View()
	println(view)
}
