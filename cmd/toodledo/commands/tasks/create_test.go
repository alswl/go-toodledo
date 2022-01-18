package tasks

import (
	"github.com/alswl/go-toodledo/pkg/models/queries"
	"reflect"
	"testing"
	"time"
)

func TestCreateWithGenerator(t *testing.T) {
	// TODO
}

func Test_cmdQuery_ToQuery(t *testing.T) {
	type fields struct {
		ContextID int64
		FolderID  int64
		GoalID    int64
		Priority  string
		DueDate   time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		want    *queries.TaskCreateQuery
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				ContextID: 1,
				FolderID:  2,
				GoalID:    3,
				Priority:  "high",
				DueDate:   time.Time{},
			},
			want: &queries.TaskCreateQuery{
				ContextID: 1,
				FolderID:  2,
				GoalID:    3,
				Priority:  2,
				DueDate:   time.Time{},
			},
			wantErr: false,
		},
		{
			name: "",
			fields: fields{
				ContextID: 0,
				FolderID:  0,
				GoalID:    0,
				Priority:  "",
				DueDate:   time.Time{},
			},
			want: &queries.TaskCreateQuery{
				Priority: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &cmdQuery{
				ContextID: tt.fields.ContextID,
				FolderID:  tt.fields.FolderID,
				GoalID:    tt.fields.GoalID,
				Priority:  tt.fields.Priority,
				DueDate:   tt.fields.DueDate,
			}
			got, err := q.ToQuery()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToQuery() got = %v, want %v", got, tt.want)
			}
		})
	}
}
