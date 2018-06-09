// Package transport is responsible for creating the http endpoints by taking in the services
// and mapping it to each endpoints
package transport

import (
	"github.com/julienschmidt/httprouter"
)

type (
	// Endpoint is an aliast to the httprouter.Handle
	Endpoint = httprouter.Handle

	// Data represents a dictionary with key string and value interfaces{}
	Data = map[string]interface{}

	// Endpoints represents the services exposed as http routes
	Endpoints interface {
		Wrap(r *httprouter.Router)
	}

	// Transport represents the http transport
	Transport interface {
		Init(endpoints ...Endpoints)
	}

	transport struct {
		router *httprouter.Router
	}
)

func (t *transport) Init(endpoints ...Endpoints) {
	for _, e := range endpoints {
		e.Wrap(t.router)
	}
}

// New creates the endpoints
func New(r *httprouter.Router) Transport {
	return &transport{r}
}
