package metrics

import "github.com/prometheus/client_golang/prometheus"

var RequestCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ssp_requests_total",
		Help: "Total number of HTTP requests to /ssp",
	},
)

var RequestDuration = prometheus.NewHistogram(
	prometheus.HistogramOpts{
		Name: "ssp_request_duration_seconds",
		Help: "Time taken to process /ssp requests",
	},
)
