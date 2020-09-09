package toodledo

import (
	"context"
	"github.com/stretchr/testify/assert"
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
			got, err := c.NewRequestWithParamsAndForm(tt.args.method, tt.args.urlStr, nil, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.NewRequestWithParamsAndForm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.NewRequestWithParamsAndForm() = %v, want %v", got, tt.want)
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

func TestClient_NewRequestWithParams(t *testing.T) {
	type fields struct {
		clientMu       sync.Mutex
		client         *http.Client
		BaseURL        *url.URL
		accessToken    string
		common         Service
		AccountService *AccountService
		FolderService  *FolderService
		GoalService    *GoalService
		TaskService    *TaskService
	}
	type args struct {
		method string
		urlStr string
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Request
		wantErr bool
	}{
		{
			name:   "",
			fields: fields{},
			args: args{
				method: "POST",
				urlStr: "http://test.com",
				params: map[string]string{"a": "b"},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				clientMu:       tt.fields.clientMu,
				client:         tt.fields.client,
				BaseURL:        tt.fields.BaseURL,
				accessToken:    tt.fields.accessToken,
				common:         tt.fields.common,
				AccountService: tt.fields.AccountService,
				FolderService:  tt.fields.FolderService,
				GoalService:    tt.fields.GoalService,
				TaskService:    tt.fields.TaskService,
			}
			got, err := c.NewRequestWithParams(tt.args.method, tt.args.urlStr, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRequestWithParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, "a=b", got.URL.RawQuery)
		})
	}
}
