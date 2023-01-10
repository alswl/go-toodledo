package time

import (
	"strings"
	"time"
)

func ParseTimeStampToDuration(ts int64) time.Duration {
	now := time.Now()
	return now.Sub(time.Unix(ts, 0))
}

func ParseDurationToReadable(duration time.Duration) string {
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
