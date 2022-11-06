package services

import (
	"encoding/json"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnCompleteMarshal(t *testing.T) {
	task := models.TaskEdit{Completed: utils.WrapPointerInt64(0)}
	bytes, _ := json.Marshal([]models.TaskEdit{task})
	assert.Equal(t, "[{\"completed\":0}]", string(bytes))
}

func TestTaskEditOnlyOneField(t *testing.T) {
	task := models.Task{Title: "new"}
	bytes, _ := json.Marshal(task)
	assert.Equal(t, "{\"title\":\"new\"}", string(bytes))
}
