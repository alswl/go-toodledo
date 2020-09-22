package toodledo

import (
	"bytes"
	"encoding/json"
	"github.com/alswl/go-toodledo/pkg/toodledo/models"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestGoalService_Add(t *testing.T) {
	var goal *models.Goal
	input := "{\"errorCode\":401,\"errorDesc\":\"Your goal must have a name\"}"

	reader := ioutil.NopCloser(bytes.NewBuffer([]byte(input)))
	decErr := json.NewDecoder(reader).Decode(goal)
	assert.NotNil(t, decErr)
}
