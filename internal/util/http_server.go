package util

import (
	"net/http"
	"time"
)

// NewHTTPServer returns a pointer to http.Server with pre-configured timeouts
// to avoid Slowloris attack
func NewHTTPServer(addr string, r http.Handler) *http.Server {
	return &http.Server{
		Addr:           addr, // ":" + addr
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
