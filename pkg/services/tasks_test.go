package services

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/alswl/go-common/pointers"

	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestUnCompleteMarshal(t *testing.T) {
	task := models.TaskEdit{Completed: pointers.WrapPointerInt64(0)}
	bytes, _ := json.Marshal([]models.TaskEdit{task})
	assert.Equal(t, `[{"completed":0}]`, string(bytes))
}

func TestTaskEditOnlyOneField(t *testing.T) {
	task := models.TaskEdit{Title: pointers.WrapPointerString("new")}
	bytes, _ := json.Marshal(task)
	assert.Equal(t, "{\"title\":\"new\"}", string(bytes))
}

func Test_sortTasks(t *testing.T) {
	sorted := sortTasks([]*models.Task{
		{Priority: 1, Title: "a"},
		{Priority: 2, Title: "b"},
		{Priority: 3, Title: "c"},
	})
	assert.Equal(t, "c", sorted[0].Title)
	assert.Equal(t, "b", sorted[1].Title)
	assert.Equal(t, "a", sorted[2].Title)

	sorted2 := sortTasks([]*models.Task{
		{Priority: 1, Title: "a"},
		{Priority: 2, Status: 5, Title: "b"},
		{Priority: 2, Status: 6, Title: "c"},
	})
	assert.Equal(t, "c", sorted2[0].Title)
	assert.Equal(t, "b", sorted2[1].Title)
	assert.Equal(t, "a", sorted2[2].Title)

	sorted3 := sortTasks([]*models.Task{
		{Priority: 1, Title: "a"},
		{Priority: 2, Status: 5, Title: "b"},
		{Priority: 2, Status: 6, Title: "c"},
		{Priority: 2, Status: 6, Duedate: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC).Unix(), Title: "d"},
		{Priority: 2, Status: 6, Duedate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC).Unix(), Title: "e"},
	})
	assert.Equal(t, "e", sorted3[0].Title)
	assert.Equal(t, "d", sorted3[1].Title)
	assert.Equal(t, "c", sorted3[2].Title)
	assert.Equal(t, "b", sorted3[3].Title)
	assert.Equal(t, "a", sorted3[4].Title)
}

func Test_rankTask(t *testing.T) {
	assert.Equal(t, int64(300000), rankTask(&models.Task{Priority: 1}))
	assert.Equal(t, int64(400000), rankTask(&models.Task{Priority: 2}))
}
