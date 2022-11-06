package models

import (
	"github.com/alswl/go-toodledo/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/thoas/go-funk"
	"reflect"
	"testing"
	"time"
)

func TestRichTask_DueDate(t1 *testing.T) {
	type fields struct {
		Task       Task
		TheContext Context
		TheFolder  Folder
		TheGoal    Goal
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{
			name: "",
			fields: fields{
				Task:       Task{Duedate: 1645704000},
				TheContext: Context{},
				TheFolder:  Folder{},
				TheGoal:    Goal{},
			},
			want: time.Date(2022, 02, 24, 20, 0, 0, 0, utils.ChinaTimeZone),
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := RichTask{
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
		Task       Task
		TheContext Context
		TheFolder  Folder
		TheGoal    Goal
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "",
			fields: fields{
				Task:       Task{Duedate: 1645704000, Duetime: 2*60*60 + 1*60},
				TheContext: Context{},
				TheFolder:  Folder{},
				TheGoal:    Goal{},
			},
			want: "2022-02-24 02:01",
		},
		{
			name: "",
			fields: fields{
				Task:       Task{},
				TheContext: Context{},
				TheFolder:  Folder{},
				TheGoal:    Goal{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := RichTask{
				Task:       tt.fields.Task,
				TheContext: tt.fields.TheContext,
				TheFolder:  tt.fields.TheFolder,
				TheGoal:    tt.fields.TheGoal,
			}
			if got := t.DueString(); got != tt.want {
				t1.Errorf("DueString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRIchTask_Timer(t *testing.T) {
	var task = RichTask{}
	task.Timer = 5
	task.Timeron = time.Date(2022, 11, 06, 13, 00, 00, 0, utils.ChinaTimeZone).Unix()
	assert.NotNil(t, task.TimerString())
}

func TestFunkIsEmpty(t *testing.T) {
	var pointerNil *int = nil
	assert.True(t, funk.IsEmpty(pointerNil))
	var ten int = int(10)
	var pointerTen = &ten
	assert.False(t, funk.IsEmpty(pointerTen))
	var zero int = int(0)
	var pointerZero = &zero
	assert.True(t, funk.IsEmpty(pointerZero))
}

func TestFunkIsZero(t *testing.T) {
	var pointerNil *int = nil
	assert.True(t, funk.IsZero(pointerNil))
	var ten int = int(10)
	var pointerTen = &ten
	assert.False(t, funk.IsZero(pointerTen))
	var zero int = int(0)
	var pointerZero = &zero
	// NOTICE, 0 is not zero
	assert.False(t, funk.IsZero(pointerZero))
}
