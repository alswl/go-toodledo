package time

import (
	"testing"
	"time"
)

func TestParseDurationToReadable(t *testing.T) {
	type args struct {
		duration time.Duration
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{duration: time.Second * 10},
			want: "10s",
		},
		{
			name: "",
			args: args{duration: time.Hour * 10},
			want: "10h",
		},
		{
			name: "",
			args: args{duration: time.Hour * 24 * 10},
			want: "168h",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseDurationToReadable(tt.args.duration); got != tt.want {
				t.Errorf("ParseDurationToReadable() = %v, want %v", got, tt.want)
			}
		})
	}
}
