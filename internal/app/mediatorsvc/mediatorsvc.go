// Package mediatorsvc handles multiple service orchestrations
// and is categorized as domain model
package mediatorsvc

// New returns a new mediator service
func New(m Mediator, middlewares ...Middleware) Service {
	s := NewService(m)
	s = Decorate(s, middlewares...)
	return s
}
