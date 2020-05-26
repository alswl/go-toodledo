package toodledo

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"sync"
	"testing"
)

func TestClient_NewRequest(t *testing.T) {
	type fields struct {
		clientMu sync.Mutex
		client   *http.Client
		BaseURL  *url.URL
		Folder   *FolderService
	}
	type args struct {
		method string
		urlStr string
		body   interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Request
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				clientMu:      tt.fields.clientMu,
				client:        tt.fields.client,
				BaseURL:       tt.fields.BaseURL,
				FolderService: tt.fields.Folder,
			}
			got, err := c.NewRequest(tt.args.method, tt.args.urlStr, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.NewRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.NewRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Do(t *testing.T) {
	type fields struct {
		clientMu sync.Mutex
		client   *http.Client
		BaseURL  *url.URL
		Folder   *FolderService
	}
	type args struct {
		ctx context.Context
		req *http.Request
		v   interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				clientMu:      tt.fields.clientMu,
				client:        tt.fields.client,
				BaseURL:       tt.fields.BaseURL,
				FolderService: tt.fields.Folder,
			}
			got, err := c.Do(tt.args.ctx, tt.args.req, tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		httpClient *http.Client
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(""); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
