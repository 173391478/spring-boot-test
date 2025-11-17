
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

func (m *MetricsCollector) RecordRequest(latency time.Duration, statusCode int) {
    m.requestCount++
    m.totalLatency += latency
    
    if statusCode >= 400 {
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
    return float64(m.errorCount) / float64(m.requestCount) * 100
}

func (m *MetricsCollector) Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
        next.ServeHTTP(rw, r)
        
        latency := time.Since(start)
        m.RecordRequest(latency, rw.statusCode)
    })
}

type responseWriter struct {
    http.ResponseWriter
    statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
    rw.statusCode = code
    rw.ResponseWriter.WriteHeader(code)
}

func main() {
    collector := NewMetricsCollector()
    
    mux := http.NewServeMux()
    mux.Handle("/", collector.Middleware(http.HandlerFunc(handler)))
    mux.Handle("/metrics", collector.Middleware(http.HandlerFunc(metricsHandler(collector))))
    
    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}

func handler(w http.ResponseWriter, r *http.Request) {
    time.Sleep(10 * time.Millisecond)
    
    if r.URL.Path == "/error" {
        http.Error(w, "Not Found", http.StatusNotFound)
        return
    }
    
    fmt.Fprintf(w, "Hello, World!")
}

func metricsHandler(collector *MetricsCollector) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        fmt.Fprintf(w, `{
            "total_requests": %d,
            "average_latency_ms": %.2f,
            "error_rate_percent": %.2f
        }`, 
            collector.requestCount,
            float64(collector.GetAverageLatency().Microseconds())/1000.0,
            collector.GetErrorRate())
    }
}