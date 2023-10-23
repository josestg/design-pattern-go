package adapter

import (
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"
)

// HealthHandler is a handler for health check.
type HealthHandler struct {
	log *slog.Logger
}

// NewHealthHandler creates a new health handler.
func NewHealthHandler(log *slog.Logger) *HealthHandler {
	return &HealthHandler{log: log}
}

// ServeHTTP implements http.Handler.
func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	started := time.Now()
	l := h.log.With("method", r.Method, "uri", r.RequestURI)

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

	health := NewHealthHandler(log)
	mux.Handle("/api/v1/health", health)
}
