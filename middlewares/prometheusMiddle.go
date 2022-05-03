package middlewares

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// MetricsMiddleware ..
type MetricsMiddleware struct {
	opsProcessed       *prometheus.CounterVec
	histogramProcessed *prometheus.HistogramVec
}

// NewMetricsMiddleware ..
func NewMetricsMiddleware() *MetricsMiddleware {

	// Counter
	opsProcessed := promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "myapp_request_total",
		Help: "The total number of processed events",
	}, []string{"method", "route", "statuscode"})

	// Histogram
	buckets := []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10}
	responseTimeHistogram := promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "myapp_http_server_request_duration_seconds",
		Help:    "Histogram of response time for handler in seconds",
		Buckets: buckets,
	}, []string{"method", "route"})

	return &MetricsMiddleware{
		opsProcessed:       opsProcessed,
		histogramProcessed: responseTimeHistogram,
	}
}

// Metrics middleware to collect metrics from http requests
func (lm *MetricsMiddleware) Metrics(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// Set histogram value
		timer := prometheus.NewTimer(lm.histogramProcessed.With(
			prometheus.Labels{
				"method": c.Request().Method,
				"route":  c.Request().RequestURI,
			}))
		defer timer.ObserveDuration()

		// Run main handler
		if err := next(c); err != nil {
			c.Error(err)
		}

		// Add value for count request
		lm.opsProcessed.With(
			prometheus.Labels{
				"method":     c.Request().Method,
				"route":      c.Request().RequestURI,
				"statuscode": strconv.Itoa(c.Response().Status)}).Inc()

		return nil
	}
}
