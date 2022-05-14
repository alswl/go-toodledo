package tasks

import (
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	mockservices "github.com/alswl/go-toodledo/test/mock/services"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

func Test_cmdQuery_ToQuery(t *testing.T) {
	contextSvc := &mockservices.ContextService{}
	contextSvc.On("Find", mock.Anything).Return(&models.Context{ID: 1}, nil)
	folderSvc := &mockservices.FolderService{}
	folderSvc.On("Find", mock.Anything).Return(&models.Folder{ID: 2}, nil)
	goalSvc := &mockservices.GoalService{}
	goalSvc.On("Find", mock.Anything).Return(&models.Goal{ID: 3}, nil)

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
			got, err := q.ToQuery(contextSvc, folderSvc, goalSvc)
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
