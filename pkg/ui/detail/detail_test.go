package detail

import (
	"testing"

	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestView(t *testing.T) {
	task := models.RichTask{
		Task: models.Task{
			Title: "test",
			Note:  "nnn\nyyy",
		},
		TheContext: &models.Context{Name: "c"},
		TheFolder:  &models.Folder{Name: "f"},
		TheGoal:    &models.Goal{Name: "g"},
	}
	m := New(task)
	m.Resize(400, 20)
	view := m.View()
	assert.Equal(t,
		`┌────────────────────────────────────────────────────────────────────┐
│Link:               https://www.toodledo.com/tasks/index.php?#task_0│
│Completed:          [ ]                                             │
│ID:                 0                                               │
│Title:              test                                            │
│Context:            c                                               │
│Folder:             f                                               │
│Goal:               g                                               │
│Status:             None                                            │
│Due:                                                                │
│Repeat:                                                             │
│Priority:           Low                                             │
│Repeat:                                                             │
│Length:                                                             │
│Timer:                                                              │
│Tag:                                                                │
│Star:               false                                           │
│Note:               nnn                                             │
└────────────────────────────────────────────────────────────────────┘`,
		view,
	)
}

func TestViewMinimal(t *testing.T) {
	task := models.RichTask{
		Task: models.Task{
			Title: "test",
			Note:  "nnn\nyyy",
		},
		TheContext: &models.Context{Name: "c"},
		TheFolder:  &models.Folder{Name: "f"},
		TheGoal:    &models.Goal{Name: "g"},
	}
	m := New(task)
	m.Resize(400, 5)
	view := m.View()
	assert.Equal(t,
		`┌────────────────────────────────────────────────────────────────────┐
│Link:               https://www.toodledo.com/tasks/index.php?#task_0│
│Completed:          [ ]                                             │
└────────────────────────────────────────────────────────────────────┘`,
		view,
	)
}
