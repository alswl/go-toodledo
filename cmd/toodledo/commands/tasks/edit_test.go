package tasks

import (
	"github.com/stretchr/testify/assert"
	"github.com/thoas/go-funk"
	"testing"
)

func TestIsZero(t *testing.T) {
	q := cmdEditQuery{}
	assert.True(t, funk.IsZero(q))
	assert.False(t, funk.IsZero(&q))

	q = cmdEditQuery{
		Context:  "",
		Folder:   "",
		Goal:     "",
		Priority: "",
		Status:   "",
		DueDate:  "",
		Title:    "",
	}
	assert.True(t, funk.IsZero(q))
	assert.False(t, funk.IsZero(&q))
}
