package toodledo

import (
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"bytes"
)

func TestGoalService_Add(t *testing.T) {
	var goal *Goal
	input := "{\"errorCode\":401,\"errorDesc\":\"Your goal must have a name\"}"

	reader := ioutil.NopCloser(bytes.NewBuffer([]byte(input)))
	decErr := json.NewDecoder(reader).Decode(goal)
	assert.NotNil(t, decErr)
}
