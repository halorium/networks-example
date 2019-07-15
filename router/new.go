package router

import "github.com/halorium/httprouter"

// Router is our internal router
type Router struct {
	*httprouter.Router
}

// New creates a new Router
func New() *Router {
	router := httprouter.New()

	router.UseRawPath = true

	router.PanicHandler = panicHandler

	router.NotFound = &notFoundHandler{}

	return &Router{
		router,
	}
}
