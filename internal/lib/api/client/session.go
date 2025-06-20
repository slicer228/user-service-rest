package client

import (
	"net/http"
	"time"
)

func NewHttpClientSession() *http.Client {
	return &http.Client{
		Timeout: 5 * time.Second,
	}
}
