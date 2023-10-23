package adapter

import (
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"
)

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

	mux.HandleFunc("/api/v1/health", HealthCheckHandler)
}
