
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

type MetricsCollector struct {
    requestCount    int
    totalLatency   time.Duration
    errorCount     int
}

func NewMetricsCollector() *MetricsCollector {
    return &MetricsCollector{}
}

func (m *MetricsCollector) RecordRequest(latency time.Duration, isError bool) {
    m.requestCount++
    m.totalLatency += latency
    if isError {
        m.errorCount++
    }
}

func (m *MetricsCollector) GetAverageLatency() time.Duration {
    if m.requestCount == 0 {
        return 0
    }
    return m.totalLatency / time.Duration(m.requestCount)
}

func (m *MetricsCollector) GetErrorRate() float64 {
    if m.requestCount == 0 {
        return 0.0
    }
    return float64(m.errorCount) / float64(m.requestCount)
}

func (m *MetricsCollector) Reset() {
    m.requestCount = 0
    m.totalLatency = 0
    m.errorCount = 0
}

func (m *MetricsCollector) String() string {
    return fmt.Sprintf("Requests: %d, Avg Latency: %v, Error Rate: %.2f%%",
        m.requestCount, m.GetAverageLatency(), m.GetErrorRate()*100)
}

func main() {
    collector := NewMetricsCollector()

    http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, collector.String())
    })

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        defer func() {
            latency := time.Since(start)
            isError := false
            
            if r.URL.Path == "/error" {
                isError = true
            }
            
            collector.RecordRequest(latency, isError)
        }()

        switch r.URL.Path {
        case "/error":
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        default:
            fmt.Fprintf(w, "Hello, World!")
        }
    })

    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}