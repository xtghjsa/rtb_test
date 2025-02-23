package metrics

import "time"

func MeasureResponseTime(next func()) {
	start := time.Now()
	next()
	duration := time.Since(start)
	RequestDuration.Observe(duration.Seconds())
}
