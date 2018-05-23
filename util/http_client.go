package util

import (
	"net/http"
	"time"
)

const (
	maxIdleConnections int = 20
	idleConnTimeout    int = 5
	requestTimeout     int = 5
)

// NewHTTPClient returns a new http client with pre-configured timeouts
func NewHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:    maxIdleConnections,
			IdleConnTimeout: time.Duration(idleConnTimeout) * time.Second,
		},
		Timeout: time.Duration(requestTimeout) * time.Second,
	}
}
