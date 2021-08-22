package toodledo

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	type args struct {
		httpClient *http.Client
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		// TODO: AddExtend test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(""); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
