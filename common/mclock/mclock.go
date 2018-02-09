package mclock

import (
	"time"

	"github.com/aristanetworks/goarista/monotime"
)

// AbsTime is absolute monotonic time
type AbsTime time.Duration

// Now return a monotime form
func Now() AbsTime {
	return AbsTime(monotime.Now())
}
