package utils_test

import (
	"testing"

	"github.com/alswl/go-toodledo/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func TestWrapListPointer(t *testing.T) {
	type args struct {
		list []string
	}
	s := "a"
	s2 := "b"
	s3 := "c"
	tests := []struct {
		name string
		args args
		want []*string
	}{
		{
			name: "",
			args: args{[]string{s, s2, s3}},
			want: []*string{&s, &s2, &s3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, utils.WrapListPointer(tt.args.list), "WrapListPointer(%v)", tt.args.list)
		})
	}
}

func TestUnwrapListPointer(t *testing.T) {
	type args struct {
		list []*string
	}
	s := "a"
	s2 := "b"
	s3 := "c"
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "",
			args: args{[]*string{&s, &s2, &s3}},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, utils.UnwrapListPointer(tt.args.list), "UnwrapListPointer(%v)", tt.args.list)
		})
	}
}
