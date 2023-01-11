package models_test

import (
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/alswl/go-toodledo/pkg/models"

	"github.com/alswl/go-toodledo/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/thoas/go-funk"
)

func TestRichTask_DueDate(t1 *testing.T) {
	type fields struct {
		Task       models.Task
		TheContext *models.Context
		TheFolder  *models.Folder
		TheGoal    *models.Goal
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{
			name: "",
			fields: fields{
				Task:       models.Task{Duedate: 1645704000},
				TheContext: &models.Context{},
				TheFolder:  &models.Folder{},
				TheGoal:    &models.Goal{},
			},
			want: time.Date(2022, 02, 24, 20, 0, 0, 0, utils.ChinaTimeZone),
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := models.RichTask{
				Task:       tt.fields.Task,
				TheContext: tt.fields.TheContext,
				TheFolder:  tt.fields.TheFolder,
				TheGoal:    tt.fields.TheGoal,
			}
			if got := t.TheDueDate(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("DueDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRichTask_Due(t1 *testing.T) {
	type fields struct {
		Task       models.Task
		TheContext models.Context
		TheFolder  models.Folder
		TheGoal    models.Goal
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "",
			fields: fields{
				Task:       models.Task{Duedate: 1645704000, Duetime: 2*60*60 + 1*60},
				TheContext: models.Context{},
				TheFolder:  models.Folder{},
				TheGoal:    models.Goal{},
			},
			want: "2022-02-24 02:01",
		},
		{
			name: "",
			fields: fields{
				Task:       models.Task{},
				TheContext: models.Context{},
				TheFolder:  models.Folder{},
				TheGoal:    models.Goal{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := models.RichTask{
				Task:       tt.fields.Task,
				TheContext: &tt.fields.TheContext,
				TheFolder:  &tt.fields.TheFolder,
				TheGoal:    &tt.fields.TheGoal,
			}
			if got := t.DueString(); got != tt.want {
				t1.Errorf("DueString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRIchTask_Timer(t *testing.T) {
	var task = models.RichTask{}
	task.Timer = 5
	task.Timeron = time.Date(2022, 11, 06, 13, 00, 00, 0, utils.ChinaTimeZone).Unix()
	t.Log(task.TimerString())
	assert.NotNil(t, task.TimerString())
	assert.Contains(t, task.TimerString(), "*")
	assert.False(t, strings.HasSuffix(task.TimerString(), "s"))
}

func TestRichTask_TimerNotRunning(t *testing.T) {
	var task = models.RichTask{}
	task.Timer = 100
	task.Timeron = 0
	assert.Equal(t, "2m", task.TimerString())
}

func TestFunkIsEmpty(t *testing.T) {
	var pointerNil *int
	assert.True(t, funk.IsEmpty(pointerNil))
	var ten = 10
	var pointerTen = &ten
	assert.False(t, funk.IsEmpty(pointerTen))
	var zero = 0
	var pointerZero = &zero
	assert.True(t, funk.IsEmpty(pointerZero))
}

func TestFunkIsZero(t *testing.T) {
	var pointerNil *int
	assert.True(t, funk.IsZero(pointerNil))
	var ten = 10
	var pointerTen = &ten
	assert.False(t, funk.IsZero(pointerTen))
	var zero = 0
	var pointerZero = &zero
	// NOTICE, 0 is not zero
	assert.False(t, funk.IsZero(pointerZero))
}

func TestRichTask_LengthString(t *testing.T) {
	var m models.RichTask
	assert.Equal(t, "", m.LengthString())

	m.Length = 1
	assert.Equal(t, "1m", m.LengthString())

	m.Length = 30
	assert.Equal(t, "30m", m.LengthString())
}
