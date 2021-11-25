//go:build integration
// +build integration

package dao

import (
	"github.com/alswl/go-toodledo/test/suites/itinjector"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_cache_ListAll(t *testing.T) {
	app, err := itinjector.InitApp()
	assert.Nil(t, err)
	assert.NotNil(t, app)
	all, err := app.FolderSvc.ListAll()
	assert.Nil(t, err)
	assert.NotEmpty(t, all)
}
