package models

import (
	"github.com/alswl/go-toodledo/pkg/utils"
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
