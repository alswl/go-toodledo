package queries

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTaskCreateQueryMarshal(t *testing.T) {
	q := TaskCreateQuery{}
	bs, err := json.Marshal(q.ToModel())

	assert.NoError(t, err)
	assert.Equal(t, `{}`, string(bs))
}
