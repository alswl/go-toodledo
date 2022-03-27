package tasks

import (
	"github.com/alswl/go-toodledo/pkg/models/queries"
	mockservices "github.com/alswl/go-toodledo/test/mock/services"
	"reflect"
	"testing"
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
		DueDate   string
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
				DueDate:   "2022-01-01",
			},
			want: &queries.TaskCreateQuery{
				ContextID: 1,
				FolderID:  2,
				GoalID:    3,
				Priority:  2,
				DueDate:   "2022-01-01",
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
				DueDate:   "2022-01-01",
			},
			want: &queries.TaskCreateQuery{
				Priority: 1,
				DueDate:  "2022-01-01",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &cmdCreateQuery{
				Context:  "c-1",
				Folder:   "f-1",
				Goal:     "g-1",
				Priority: tt.fields.Priority,
				DueDate:  tt.fields.DueDate,
			}
			got, err := q.ToQuery(&mockservices.ContextService{}, &mockservices.FolderService{}, &mockservices.GoalService{})
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
