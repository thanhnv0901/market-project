package middlewares

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// MetricsMiddleware ..
type MetricsMiddleware struct {
	opsProcessed *prometheus.CounterVec
}

// NewMetricsMiddleware ..
func NewMetricsMiddleware() *MetricsMiddleware {
	opsProcessed := promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	}, []string{"method", "path", "statuscode"})
	return &MetricsMiddleware{
		opsProcessed: opsProcessed,
	}
}

// Metrics middleware to collect metrics from http requests
func (lm *MetricsMiddleware) Metrics(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		if err := next(c); err != nil {
			c.Error(err)
		}

		lm.opsProcessed.With(prometheus.Labels{"method": c.Request().Method, "path": c.Request().RequestURI, "statuscode": strconv.Itoa(c.Response().Status)}).Inc()
		return nil
	}
}
