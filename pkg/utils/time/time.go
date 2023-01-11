package time

import (
	"strings"
	"time"
)

// ParseTimestampToDuration parse ts to time stamp.
func ParseTimestampToDuration(ts int64) time.Duration {
	now := time.Now()
	return now.Sub(time.Unix(ts, 0))
}

// ParseDurationInSecondToDuration parse ts to time stamp.
func ParseDurationInSecondToDuration(second int64) time.Duration {
	return time.Duration(second) * time.Second
}

// ParseDurationToReadableShort return human-readable duration.
// it only returns main unit, so it was no accurate.
func ParseDurationToReadableShort(duration time.Duration) string {
	s := ""
	// nolint: gocritic
	if duration < time.Minute {
		s = duration.Round(time.Second).String()
	} else if duration < time.Hour {
		s = duration.Round(time.Minute).String()
	} else if duration < time.Hour*24 {
		s = duration.Round(time.Hour).String()
	} else if duration < time.Hour*24*7 {
		// nolint: gomnd
		s = duration.Round(time.Hour * 24).String()
	} else if duration < time.Hour*24*30 {
		// nolint: gomnd
		s = duration.Round(time.Hour * 24 * 7).String()
	} else if duration < time.Hour*24*365 {
		// nolint: gomnd
		s = duration.Round(time.Hour * 24 * 30).String()
	} else {
		// nolint: gomnd
		s = duration.Round(time.Hour * 24 * 365).String()
	}
	if strings.HasSuffix(s, "m0s") {
		s = s[:len(s)-2]
	}
	if strings.HasSuffix(s, "h0m") {
		s = s[:len(s)-2]
	}
	return s
}

// ParseDurationToReadable return human-readable duration.
func ParseDurationToReadable(duration time.Duration) string {
	duration = duration.Round(time.Minute)
	s := duration.String()

	if strings.HasSuffix(s, "m0s") {
		s = s[:len(s)-2]
	}
	if strings.HasSuffix(s, "h0m") {
		s = s[:len(s)-2]
	}
	return s
}
