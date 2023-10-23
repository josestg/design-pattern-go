package adapter

import (
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"
)

// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler that calls f.
type HandlerFunc func(http.ResponseWriter, *http.Request)

// quick check to verify if the HandlerFunc has the same signature as
// http.Handler.ServeHTTP.
var _ HandlerFunc = http.DefaultServeMux.ServeHTTP

// quick check to verify if the HandlerFunc implements http.Handler.
var _ http.Handler = HandlerFunc(nil)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

// HealthCheckHandler handles health check request.
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	started := time.Now()
	l := slog.Default().With("method", r.Method, "uri", r.RequestURI)

	w.WriteHeader(http.StatusOK)
	if _, err := io.WriteString(w, "OK"); err != nil {
		l.Error("could not write response", "error", err)
	} else {
		l.Info("health check success", "latency", time.Since(started))
	}
}

// RegisterRoutes registers all routes to mux.
func RegisterRoutes(mux *http.ServeMux) {
	log := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{}))
	slog.SetDefault(log)

	mux.Handle("/api/v1/health", HandlerFunc(HealthCheckHandler))
}
