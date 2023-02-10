package sidebar

import (
	"testing"

	"github.com/alswl/go-toodledo/pkg/models"

	"github.com/charmbracelet/bubbles/list"
)

func TestSidebar(t *testing.T) {
	// set CI not works
	// t.Setenv("CI", "true")
	// os.Setenv("CI", "true")
	items := []list.Item{
		models.NewItem(123, "Raspberry Pi’s"),
		models.NewItem(123, "Nutella"),
		models.NewItem(123, "Bitter melon"),
		models.NewItem(123, "Nice socks"),
		models.NewItem(123, "Eight hours of sleep"),
		models.NewItem(123, "Cats"),
		models.NewItem(123, "Plantasia, the album"),
		models.NewItem(123, "Pour over coffee"),
		models.NewItem(123, "VR"),
		models.NewItem(123, "Noguchi Lamps"),
		models.NewItem(123, "Linux"),
		models.NewItem(123, "Business school"),
		models.NewItem(123, "Pottery"),
		models.NewItem(123, "Shampoo"),
		models.NewItem(123, "Table tennis"),
		models.NewItem(123, "Milk crates"),
		models.NewItem(123, "Afternoon tea"),
		models.NewItem(123, "Stickers"),
		models.NewItem(123, "20° Weather"),
		models.NewItem(123, "Warm light"),
		models.NewItem(123, "The vernal equinox"),
		models.NewItem(123, "Gaffer’s tape"),
		models.NewItem(123, "Terrycloth"),
	}

	m := Model{
		states:      NewStates(),
		contextList: list.New(items, list.NewDefaultDelegate(), 0, 0),
		folderList:  list.New(items, list.NewDefaultDelegate(), 0, 0),
		goalList:    list.New(items, list.NewDefaultDelegate(), 0, 0),
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
