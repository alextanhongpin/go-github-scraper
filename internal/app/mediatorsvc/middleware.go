package mediatorsvc

// Middleware represents a decorator pattern
type Middleware func(Service) Service

// Decorate decorates a service with a list of provided middlewares
func Decorate(s Service, ms ...Middleware) Service {
	decorated := s
	for _, m := range ms {
		decorated = m(decorated)
	}
	return decorated
}
