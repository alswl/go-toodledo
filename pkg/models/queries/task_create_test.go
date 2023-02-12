package queries_test

import (
	"encoding/json"
	"testing"

	"github.com/alswl/go-toodledo/pkg/models/queries"

	"github.com/stretchr/testify/assert"
)

func TestTaskCreateQueryMarshal(t *testing.T) {
	q := queries.TaskCreateQuery{}
	bs, err := json.Marshal(q.ToModel())

	assert.NoError(t, err)
	// default marshall of empty task
	assert.Equal(t, `{"priority":0,"star":0,"status":0,"via":""}`, string(bs))
}
