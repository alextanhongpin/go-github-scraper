package usersvc

// Middleware represents a function that takes a service and returns the service with middleware
type Middleware func(Service) Service

// Decorate takes a service and a list of middlewares and return the decorated service
func Decorate(s Service, ms ...Middleware) Service {
	decorated := s
	for _, m := range ms {
		decorated = m(decorated)
	}
	return decorated
}
