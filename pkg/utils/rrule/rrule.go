package rrule

import "github.com/teambition/rrule-go"

// ParseToodledoRRule parses a toodledo rrule to a string
// it was only implemented for the Toodledo cases.
func ParseToodledoRRule(rule rrule.RRule) string {
	opt := rule.OrigOptions

	if opt.Interval == 0 {
		opt.Interval = 1
	}
	// simple format
	if len(opt.Byweekno) == 0 &&
		len(opt.Byyearday) == 0 &&
		len(opt.Bymonthday) == 0 &&
		len(opt.Byweekday) == 0 &&
		len(opt.Byeaster) == 0 {
		if opt.Interval == 1 {
			// nolint: exhaustive
			switch opt.Freq {
			case rrule.YEARLY:
				return "Yearly"
			case rrule.MONTHLY:
				return "Monthly"
			case rrule.WEEKLY:
				return "Weekly"
			case rrule.DAILY:
				return "Daily"
			default:
				return "..."
			}
		}
		// nolint: gomnd
		if opt.Interval == 2 {
			// nolint: exhaustive
			switch opt.Freq {
			case rrule.YEARLY:
				return "Every other year"
			case rrule.MONTHLY:
				return "Bimonthly"
			case rrule.WEEKLY:
				return "Biweekly"
			case rrule.DAILY:
				return "Every other day"
			default:
				return "..."
			}
		}
		// nolint: gomnd
		if opt.Interval == 3 {
			if opt.Freq == rrule.MONTHLY {
				return "Quarterly"
			}
		}
		// nolint: gomnd
		if opt.Interval == 6 {
			if opt.Freq == rrule.MONTHLY {
				return "Semiannually"
			}
		}
	}
	// handle monthly
	if opt.Freq == rrule.MONTHLY {
		if len(opt.Bymonthday) == 0 {
			return "Monthly"
		}
		if len(opt.Bymonthday) == 1 && opt.Bymonthday[0] == -1 {
			return "End of month"
		}
		// NOTICE, Toodledo rrule did not include a DTStart
		// and rrule using current date as DTStart, and the bymonthday will be calculated
		// so here ignore the bymonthday
		return "Monthly"
	}

	return "Custom"
}
