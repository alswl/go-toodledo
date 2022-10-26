package services

import (
	"encoding/json"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnCompleteMarshal(t *testing.T) {
	task := models.Task{Completed: 0}
	bytes, _ := json.Marshal([]models.Task{task})
	assert.Equal(t, "[{\"completed\":0}]", string(bytes))
}
