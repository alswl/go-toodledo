package utils_test

import (
	"testing"
	"time"

	"github.com/alswl/go-toodledo/pkg/utils"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
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
	err := utils.BindFlagsByQuery(&cmd, Q{})
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
	err := utils.BindFlagsByQuery(&cmd, Q{})
	assert.NoError(t, err)

	q := Q{}
	_ = cmd.Flags().Set("s", "test")
	_ = cmd.Flags().Set("i", "1")
	_ = cmd.Flags().Set("i64", "2")
	_ = cmd.Flags().Set("t", "2018-01-01")
	_ = cmd.Flags().Set("b", "true")
	_ = cmd.Flags().Set("ss", "a")
	_ = cmd.Flags().Set("ss", "b")

	err = utils.FillQueryByFlags(&cmd, &q)
	assert.NoError(t, err)
	assert.Equal(t, "test", q.S)
	assert.Equal(t, 1, q.I)
	assert.Equal(t, int64(2), q.I64)
	assert.Equal(t, time.Date(2018, 1, 1, 0, 0, 0, 0, utils.DefaultTimeZone), q.T)
	assert.Equal(t, true, q.B)
	assert.Equal(t, []string{"a", "b"}, q.SS)
}
