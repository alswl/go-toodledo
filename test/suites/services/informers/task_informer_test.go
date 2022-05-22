//go:build integration
// +build integration

package informers

import (
	"github.com/alswl/go-toodledo/pkg/services/informers"
	"github.com/alswl/go-toodledo/test/suites/itinjector"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRun(t *testing.T) {
	backend, err := itinjector.InitBackend()
	assert.NoError(t, err)

	informer := informers.NewTaskInformer(nil, backend)
	all, _, err := informer.ListAll()
	assert.Error(t, err)
}
