package editor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEditor(t *testing.T) {
	_, err := NewEditor("vi")
	assert.NoError(t, err)

	// using EDITOR
	_, err = NewEditor("foo-bar-not-exist")
	assert.NoError(t, err)

	// using default
	t.Setenv("EDITOR", "")
	_, err = NewEditor("foo-bar-not-exist")
	assert.NoError(t, err)
}
