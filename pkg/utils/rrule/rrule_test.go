package rrule

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/teambition/rrule-go"
)

func TestParseToodledoRRule(t *testing.T) {
	r, _ := rrule.StrToRRule("FREQ=YEARLY")
	assert.Equal(t, "Yearly", ParseToodledoRRule(*r))

	r, _ = rrule.StrToRRule("FREQ=MONTHLY")
	assert.Equal(t, "Monthly", ParseToodledoRRule(*r))

	r, _ = rrule.StrToRRule("FREQ=WEEKLY")
	assert.Equal(t, "Weekly", ParseToodledoRRule(*r))

	r, _ = rrule.StrToRRule("FREQ=DAILY")
	assert.Equal(t, "Daily", ParseToodledoRRule(*r))

	r, _ = rrule.StrToRRule("FREQ=YEARLY;INTERVAL=2")
	assert.Equal(t, "Every other year", ParseToodledoRRule(*r))

	r, _ = rrule.StrToRRule("FREQ=MONTHLY;INTERVAL=2")
	assert.Equal(t, "Bimonthly", ParseToodledoRRule(*r))

	r, _ = rrule.StrToRRule("FREQ=WEEKLY;INTERVAL=2")
	assert.Equal(t, "Biweekly", ParseToodledoRRule(*r))

	r, _ = rrule.StrToRRule("FREQ=DAILY;INTERVAL=2")
	assert.Equal(t, "Every other day", ParseToodledoRRule(*r))

	r, _ = rrule.StrToRRule("FREQ=MONTHLY;INTERVAL=3")
	assert.Equal(t, "Quarterly", ParseToodledoRRule(*r))

	r, _ = rrule.StrToRRule("FREQ=MONTHLY;INTERVAL=6")
	assert.Equal(t, "Semiannually", ParseToodledoRRule(*r))

	r, _ = rrule.StrToRRule("FREQ=MONTHLY;BYMONTHDAY=-1")
	assert.Equal(t, "End of month", ParseToodledoRRule(*r))

	r, _ = rrule.StrToRRule("FREQ=DAILY;BYMONTHDAY=1")
	assert.Equal(t, "Custom", ParseToodledoRRule(*r))
}

func TestMonthly(t *testing.T) {
	r, _ := rrule.StrToRRule("FREQ=MONTHLY")
	assert.Equal(t, "Monthly", ParseToodledoRRule(*r))
}
