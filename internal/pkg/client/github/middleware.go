package github

// Middleware represents the middleware func
type Middleware func(Service) Service

// Decorate represents the decorator that adds capabilities to services with middlewares
func Decorate(s Service, ms ...Middleware) Service {
	decorated := s
	for _, m := range ms {
		decorated = m(decorated)
	}
	return decorated
}
