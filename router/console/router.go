package router

import "reflect"

type Router struct {
	routes RouteCollections

	controllerCallBack func(controllerName string) reflect.Value
}

func New() *Router {
	return &Router{
		routes: RouteCollections{},
	}
}

func (router *Router) RegisterControllerCallBack(callback func(controllerName string) reflect.Value) {
	router.controllerCallBack = callback
}

func (router *Router) Command(signature string, description string) *Route {
	return router.addCommand(signature, description)
}

func (router *Router) addCommand(signature string, description string) *Route {
	newRoute := router.routes.Add(
		*router.createRoute(signature, description),
	)

	return newRoute
}

func (router Router) createRoute(signature string, description string) *Route {
	return router.newRoute(signature, description)
}

func (router Router) newRoute(signature string, description string) *Route {
	return NewRoute(signature, description, router)
}
