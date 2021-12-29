package pkg

import (
	"github.com/alswl/go-toodledo/pkg/models/enums/tasks"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type Q struct {
	s   string
	i64 int64
	i   int
	t   time.Time
	srt tasks.DueDateMode
	b   bool
}

func TestGenerateFlagsByStructure(t *testing.T) {
	cmd := cobra.Command{}
	err := GenerateFlagsByStructure(&cmd, Q{})
	assert.NoError(t, err)

	assert.Equal(t, `Usage:

Flags:
      --b          b
      --i int      i
      --i64 int    i64
      --s string   s
      --srt int    srt
      --t string   t
`, cmd.UsageString())
}
