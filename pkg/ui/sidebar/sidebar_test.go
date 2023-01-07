package sidebar

import (
	"testing"

	"github.com/charmbracelet/bubbles/list"
)

func TestSidebar(t *testing.T) {
	// set CI not works
	// t.Setenv("CI", "true")
	// os.Setenv("CI", "true")
	items := []list.Item{
		Item{title: "Raspberry Pi’s"},
		Item{title: "Nutella"},
		Item{title: "Bitter melon"},
		Item{title: "Nice socks"},
		Item{title: "Eight hours of sleep"},
		Item{title: "Cats"},
		Item{title: "Plantasia, the album"},
		Item{title: "Pour over coffee"},
		Item{title: "VR"},
		Item{title: "Noguchi Lamps"},
		Item{title: "Linux"},
		Item{title: "Business school"},
		Item{title: "Pottery"},
		Item{title: "Shampoo"},
		Item{title: "Table tennis"},
		Item{title: "Milk crates"},
		Item{title: "Afternoon tea"},
		Item{title: "Stickers"},
		Item{title: "20° Weather"},
		Item{title: "Warm light"},
		Item{title: "The vernal equinox"},
		Item{title: "Gaffer’s tape"},
		Item{title: "Terrycloth"},
	}

	m := Model{
		states:      NewStates(),
		contextList: list.New(items, list.NewDefaultDelegate(), 0, 0),
	}
	m.Resize(40, 20)
	view := m.View()
	t.Log(view)
	// TODO render view include un-printable chars
	//	assert.Equal(t, `┌──────────────────────────────────────┐
	// │<Contexts>                            │
	// │   List                               │
	// │                                      │
	// │  23 items                            │
	// │                                      │
	// │                                      │
	// │  1/23                                │
	// │                                      │
	// │  ↑/k up • ↓/j down • / filter • q    │
	// │quit • ? more                         │
	// │                                      │
	// │                                      │
	// │                                      │
	// │                                      │
	// │                                      │
	// │                                      │
	// │                                      │
	// │                                      │
	// │                                      │
	// └──────────────────────────────────────┘
	// `, view)
}
