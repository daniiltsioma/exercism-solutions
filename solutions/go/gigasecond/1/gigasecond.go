
// Package gigasecond helps calculate dates
// one gigasecond in the future.
package gigasecond

import "time"

// AddGigasecond returns a timestamp one gigasecond after the parameter.
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second * 1000000000)

	return t
}
