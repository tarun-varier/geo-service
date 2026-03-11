package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	httpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)
)

type GPSLocation struct {
	Room      string  `json:"room"`
	Building  string  `json:"building"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Accuracy  int     `json:"accuracy"`
	Timestamp string  `json:"timestamp"`
}

func getClassroomLocation() GPSLocation {
	return GPSLocation{
		Room:      "Room 301",
		Building:  "Computer Science Building",
		Latitude:  40.7128,
		Longitude: -74.0060,
		Accuracy:  10,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
}

func locationHandler(c *gin.Context) {
	start := time.Now()
	location := getClassroomLocation()
	duration := time.Since(start).Seconds()

	httpRequestsTotal.WithLabelValues(c.Request.Method, "/api/v1/location", "200").Inc()
	httpRequestDuration.WithLabelValues(c.Request.Method, "/api/v1/location").Observe(duration)

	c.JSON(http.StatusOK, location)
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
		"time":   time.Now().UTC().Format(time.RFC3339),
	})
}

func metricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start).Seconds()

		status := c.Writer.Status()
		if status > 0 {
			httpRequestsTotal.WithLabelValues(c.Request.Method, c.FullPath(), string(rune(status))).Inc()
		}
		httpRequestDuration.WithLabelValues(c.Request.Method, c.FullPath()).Observe(duration)
	}
}

func main() {
	r := gin.Default()

	r.Use(metricsMiddleware())

	r.GET("/api/v1/location", locationHandler)
	r.GET("/health", healthHandler)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"service": "Classroom GPS Service",
			"version": "1.0.0",
			"endpoints": []string{
				"/api/v1/location - Get GPS coordinates",
				"/health - Health check",
				"/metrics - Prometheus metrics",
			},
		})
	})

	port := "8080"
	println("Server starting on port " + port)
	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}
