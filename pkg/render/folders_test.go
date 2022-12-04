package render

import (
	"testing"

	"github.com/alswl/go-toodledo/pkg/models"
)

func TestTablesRender(t *testing.T) {
	type args struct {
		folders []*models.Folder
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{[]*models.Folder{{
				Archived: 0,
				ID:       0,
				Name:     "foor",
				Ord:      0,
				Private:  0,
			}}},
			want: ` # │ NAME │ ARCHIVED 
───┼──────┼──────────
 0 │ foor │        0 
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Tables4Folder(tt.args.folders); got != tt.want {
				t.Errorf("Tables4Folder() = %v, want %v", got, tt.want)
			}
		})
	}
}
