package sidebar

import (
	"github.com/alswl/go-toodledo/pkg/common/logging"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/models/constants"
	"github.com/alswl/go-toodledo/pkg/ui"
	"github.com/alswl/go-toodledo/pkg/ui/common"
	"github.com/charmbracelet/bubbles/list"
	"github.com/sirupsen/logrus"
)

var defaultTabs = []string{
	constants.Contexts,
	constants.Folders,
	constants.Goals,
	// "Priority",
	// "Tags",
	// "Search",
}

type Properties struct {
}

type States struct {
	IsCollapsed bool
	// CurrentTabIndex is the index of the defaultTabs
	CurrentTabIndex      int
	ItemIndexReadonlyMap map[string]int64

	Contexts []models.Context
	Folders  []models.Folder
	Goals    []models.Goal
}

func NewStates() *States {
	return &States{
		IsCollapsed:          false,
		CurrentTabIndex:      0,
		ItemIndexReadonlyMap: map[string]int64{},
		Contexts:             []models.Context{},
		Folders:              []models.Folder{},
		Goals:                []models.Goal{},
	}
}

type ItemChangeSubscriber func(tab string, item Item) error

type Model struct {
	ui.Focusable
	ui.Resizable

	log        logrus.FieldLogger
	properties Properties
	states     *States

	// view
	// list has states(selected)
	// TODO using wrapped list
	contextList list.Model
	folderList  list.Model
	goalList    list.Model

	// handler
	onItemChangeSubscribers []ItemChangeSubscriber
}

func InitModel(p Properties) Model {
	m := Model{
		log:                     logging.GetLogger("tt"),
		properties:              p,
		states:                  NewStates(),
		onItemChangeSubscribers: []ItemChangeSubscriber{},
		contextList:             common.NewSimpleList(),
		folderList:              common.NewSimpleList(),
		goalList:                common.NewSimpleList(),
	}
	// if len(m.list.Items()) > 0 {
	//	m.list.Select(0)
	//}
	m.Blur()
	return m
}
