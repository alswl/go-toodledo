package detail

import (
	"testing"
	"time"

	"github.com/alswl/go-toodledo/pkg/utils"

	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestView(t *testing.T) {
	task := models.RichTask{
		Task: models.Task{
			Title:   "test",
			Note:    "nnn",
			Duedate: time.Date(2018, 9, 1, 10, 24, 0, 0, utils.ChinaTimeZone).Unix(),
		},
		TheContext: &models.Context{Name: "c"},
		TheFolder:  &models.Folder{Name: "f"},
		TheGoal:    &models.Goal{Name: "g"},
	}
	m := InitModel(task, 100, 20)
	m.Resize(100, 20)
	view := m.View()
	t.Log(view)
	assert.Equal(t,
		`┌──────────────────────────────────────────────────────────────────────────────────────────────────┐
│Link:               https://www.toodledo.com/tasks/index.php?#task_0                              │
│Completed:          [ ]                                                                           │
│ID:                 0                                                                             │
│Title:              test                                                                          │
│Context:            c                                                                             │
│Folder:             f                                                                             │
│Goal:               g                                                                             │
│Status:             None                                                                          │
│Due:                2018-09-01                                                                    │
│Repeat:                                                                                           │
│Priority:           Low                                                                           │
│Length:                                                                                           │
│Timer:                                                                                            │
│Tag:                                                                                              │
│Star:               false                                                                         │
│Note:               nnn                                                                           │
│Added:              1970-01-01 08:00:00                                                           │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘`,
		view,
	)
}

func TestViewMinimal(t *testing.T) {
	task := models.RichTask{
		Task: models.Task{
			Title: "test",
			Note:  "nnn",
		},
		TheContext: &models.Context{Name: "c"},
		TheFolder:  &models.Folder{Name: "f"},
		TheGoal:    &models.Goal{Name: "g"},
	}
	m := InitModel(task, 100, 20)
	m.Resize(100, 5)
	view := m.View()
	t.Log(view)
	assert.Equal(t,
		`┌──────────────────────────────────────────────────────────────────────────────────────────────────┐
│Link:               https://www.toodledo.com/tasks/index.php?#task_0                              │
│Completed:          [ ]                                                                           │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘`,
		view,
	)
}
