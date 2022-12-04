package services_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/alswl/go-toodledo/pkg/services"

	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/models/enums/tasks/subtasksview"
	"github.com/stretchr/testify/assert"
)

func TestSortSubTasksInline(t *testing.T) {
	bs, err := os.ReadFile("./testdata/rich_tasks_1.json")
	assert.NoError(t, err)
	var tasks = make([]*models.Task, 0)

	err = json.Unmarshal(bs, &tasks)
	assert.NoError(t, err)
	assert.Len(t, tasks, 5)

	ts, err := services.SortSubTasks(tasks, subtasksview.Inline)
	assert.NoError(t, err)
	assert.Len(t, ts, 5)
	assert.Equal(t, "algorithm xy / heap", ts[0].Title)
	assert.Equal(t, "五天学会绘画", ts[1].Title)
	assert.Equal(t, "晒被子", ts[2].Title)
	assert.Equal(t, "algorithm xy / queue seq", ts[3].Title)
	assert.Equal(t, "algorithm xy /", ts[4].Title)
}

func TestSortSubTasksHidden(t *testing.T) {
	bs, err := os.ReadFile("./testdata/rich_tasks_1.json")
	assert.NoError(t, err)
	var tasks = make([]*models.Task, 0)

	err = json.Unmarshal(bs, &tasks)
	assert.NoError(t, err)
	assert.Len(t, tasks, 5)

	ts, err := services.SortSubTasks(tasks, subtasksview.Hidden)
	assert.NoError(t, err)
	assert.Len(t, ts, 3)
}

func TestSortSubTasksIndented(t *testing.T) {
	bs, err := os.ReadFile("./testdata/rich_tasks_1.json")
	assert.NoError(t, err)
	var tasks = make([]*models.Task, 0)

	err = json.Unmarshal(bs, &tasks)
	assert.NoError(t, err)
	assert.Len(t, tasks, 5)

	ts, err := services.SortSubTasks(tasks, subtasksview.Indented)
	assert.NoError(t, err)
	assert.Len(t, ts, 5)

	assert.Equal(t, "五天学会绘画", ts[0].Title)
	assert.Equal(t, "晒被子", ts[1].Title)
	assert.Equal(t, "algorithm xy /", ts[2].Title)
	assert.Equal(t, "algorithm xy / heap", ts[3].Title)
	assert.Equal(t, "algorithm xy / queue seq", ts[4].Title)
}
