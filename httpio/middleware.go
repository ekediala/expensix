package httpio

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type TraceKey struct{}

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rr *responseRecorder) WriteHeader(statusCode int) {
	rr.statusCode = statusCode
	rr.ResponseWriter.WriteHeader(statusCode)
}

func TraceMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := SetTraceID(r.Context())
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func LoggingMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rr := &responseRecorder{ResponseWriter: w}
		next.ServeHTTP(rr, r)
		slog.InfoContext(r.Context(), "request",
			"url", r.URL,
			"method", r.Method,
			"took", time.Since(start),
			"statusCode", rr.statusCode,
			"ip", ReadUserIP(r),
		)
	}
}

func CORSMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "") // You might want to make this more restrictive
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Max-Age", "3600")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	}
}

type LogHandler struct {
	slog.Handler
}

func NewLogHandler(h slog.Handler) *LogHandler {
	l := LogHandler{Handler: h}
	return &l
}

func (l *LogHandler) Handle(ctx context.Context, r slog.Record) error {
	if id, ok := GetTraceID(ctx); ok {
		r.Add("trace_id", id)
	}
	return l.Handler.Handle(ctx, r)
}

func (l *LogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &LogHandler{Handler: l.Handler.WithAttrs(attrs)}
}

func (l *LogHandler) WithGroup(name string) slog.Handler {
	return &LogHandler{Handler: l.Handler.WithGroup(name)}
}

func (l *LogHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return true
}

func SetTraceID(ctx context.Context) context.Context {
	c := context.WithValue(ctx, TraceKey{}, uuid.NewString())
	return c
}

func GetTraceID(ctx context.Context) (string, bool) {
	value, ok := ctx.Value(TraceKey{}).(string)
	return value, ok
}

func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}
