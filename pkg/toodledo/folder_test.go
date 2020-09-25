package toodledo

import (
	"context"
	"github.com/alswl/go-toodledo/pkg/toodledo/models"
	"reflect"
	"testing"
)

func TestFolderService_Get(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       FolderService
		args    args
		want    []*models.Folder
		want1   *Response
		wantErr bool
	}{
		// TODO: AddExtend test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.s.Get(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
