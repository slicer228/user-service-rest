package client

import (
	"log/slog"
	"net/http"
	"time"
)

func NewHttpClientSession(log *slog.Logger) *http.Client {
	log.Debug("Creating new HTTP client session...")
	return &http.Client{
		Timeout: 5 * time.Second,
	}
}
