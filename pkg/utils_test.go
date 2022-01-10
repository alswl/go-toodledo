package pkg

import (
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type Q struct {
	S   string
	I64 int64
	I   int
	T   time.Time
	B   bool
	SS  []string
}

func TestGenerateFlagsByQuery(t *testing.T) {
	cmd := cobra.Command{}
	err := GenerateFlagsByQuery(&cmd, Q{})
	assert.NoError(t, err)

	assert.Equal(t, `Usage:

Flags:
      --b            b
      --i int        i
      --i64 int      i64
      --s string     s
      --ss strings   ss
      --t string     t
`, cmd.UsageString())
}

func TestFillQueryByFlags(t *testing.T) {
	cmd := cobra.Command{}
	err := GenerateFlagsByQuery(&cmd, Q{})
	assert.NoError(t, err)

	q := Q{}
	cmd.Flags().Set("s", "test")
	cmd.Flags().Set("i", "1")
	cmd.Flags().Set("i64", "2")
	cmd.Flags().Set("t", "2018-01-01")
	cmd.Flags().Set("b", "true")
	cmd.Flags().Set("ss", "a")
	cmd.Flags().Set("ss", "b")

	err = FillQueryByFlags(&cmd, &q)
	assert.NoError(t, err)
	assert.Equal(t, "test", q.S)
	assert.Equal(t, 1, q.I)
	assert.Equal(t, int64(2), q.I64)
	assert.Equal(t, time.Date(2018, 1, 1, 0, 0, 0, 0, time.Local), q.T)
	assert.Equal(t, true, q.B)
	assert.Equal(t, []string{"a", "b"}, q.SS)
}
