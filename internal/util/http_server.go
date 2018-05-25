package util

import (
	"net/http"
	"time"
)

const (
	readTimeout  = 10
	writeTimeout = 10
)

// NewHTTPServer returns a pointer to http.Server with pre-configured timeouts
func NewHTTPServer(addr string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:           addr,
		Handler:        handler,
		ReadTimeout:    time.Duration(readTimeout) * time.Second,
		WriteTimeout:   time.Duration(writeTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
