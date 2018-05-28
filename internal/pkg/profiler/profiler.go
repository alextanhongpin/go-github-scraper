// Package profiler contains the logic to integrate profilling with julienschmidt
// httprouter package
package profiler

import (
	"net/http"
	"net/http/pprof"

	"github.com/julienschmidt/httprouter"
)

// MakeHTTP will setup the profile
func MakeHTTP(enable bool, r *httprouter.Router) {
	if !enable {
		return
	}
	// Setup profiling
	r.HandlerFunc(http.MethodGet, "/debug/pprof/", pprof.Index)
	r.HandlerFunc(http.MethodGet, "/debug/pprof/cmdline", pprof.Cmdline)
	r.HandlerFunc(http.MethodGet, "/debug/pprof/profile", pprof.Profile)
	r.HandlerFunc(http.MethodGet, "/debug/pprof/symbol", pprof.Symbol)
	r.HandlerFunc(http.MethodGet, "/debug/pprof/trace", pprof.Trace)
	r.Handler(http.MethodGet, "/debug/pprof/goroutine", pprof.Handler("goroutine"))
	r.Handler(http.MethodGet, "/debug/pprof/heap", pprof.Handler("heap"))
	r.Handler(http.MethodGet, "/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	r.Handler(http.MethodGet, "/debug/pprof/block", pprof.Handler("block"))
}
