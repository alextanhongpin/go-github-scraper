package statsvc

// Middleware takes a service and return the service with new capabilities
type Middleware func(Service) Service

// Decorate decorates a service with the given list of middlewares
func Decorate(s Service, ms ...Middleware) Service {
	decorated := s
	for _, m := range ms {
		decorated = m(s)
	}
	return decorated
}
