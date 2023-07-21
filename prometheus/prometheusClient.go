package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// 定义指标对象
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "status"},
	)
	responseTime = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_response_time_seconds",
			Help: "HTTP response time distribution.",
			Buckets: []float64{
				0.01, 0.05, 0.1, 0.5, 1, 5, 10,
			},
		},
		[]string{"method", "status"},
	)
)

func main() {
	// 注册指标对象
	prometheus.MustRegister(requestsTotal)
	prometheus.MustRegister(responseTime)

	// 定义HTTP处理函数
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// 模拟处理时间
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		// 记录请求总数指标
		requestsTotal.With(prometheus.Labels{
			"method": r.Method,
			"status": "200",
		}).Inc()

		// 记录响应时间指标
		responseTime.With(prometheus.Labels{
			"method": r.Method,
			"status": "200",
		}).Observe(time.Since(start).Seconds())

		// 响应
		fmt.Fprintln(w, "Hello, world!")
	})

	// 暴露指标
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
