package reposvc

// Middleware utilises the decorator pattern to add new functionality to the reposvc
type Middleware func(Service) Service

// Decorate will decorate the service with the given middlewares and returns the decorated Service
func Decorate(s Service, ms ...Middleware) Service {
	decorated := s
	for _, decorator := range ms {
		decorated = decorator(decorated)
	}
	return decorated
}
