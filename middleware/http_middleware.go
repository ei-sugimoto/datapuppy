package middleware

import (
	"net/http"
	"time"

	"github.com/ei-sugimoto/datapuppy/internal"
)

type StatusRecorder struct {
	http.ResponseWriter
	StatusCode int
}

func (r *StatusRecorder) WriteHeader(code int) {
	r.StatusCode = code
	r.ResponseWriter.WriteHeader(code)
}

func NewStatusRecorder(w http.ResponseWriter) *StatusRecorder {
	// デフォルトのステータスコードは200
	return &StatusRecorder{w, http.StatusOK}
}

func DatapupyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := internal.NewLogger()
		detail := internal.Detail{
			ID:        "",
			Time:      time.Now().Format(time.RFC3339),
			RemoteIP:  r.RemoteAddr,
			Host:      r.Host,
			Method:    r.Method,
			URI:       r.RequestURI,
			UserAgent: r.UserAgent(),
			Status:    200,
			Latency:   0,
		}
		start := time.Now()
		recoder := NewStatusRecorder(w)
		next.ServeHTTP(recoder, r)
		detail.Latency = int(time.Since(start).Milliseconds())
		detail.Status = recoder.StatusCode
		logger.UpsertLogToFile(&detail)

	})

}
