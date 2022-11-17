package router

import (
	"reflect"
)

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

func (router *Router) Command(command string, description string) *Route {
	return router.addCommand(command, description)
}

func (router *Router) addCommand(command string, description string) *Route {
	return router.routes.Add(
		router.createRoute(command, description),
	)
}

func (router Router) createRoute(command string, description string) *Route {
	return router.newRoute(command, description)
}

func (router Router) newRoute(command string, description string) *Route {
	return NewRoute(command, description)
}

func (router Router) GetRoutes() RouteCollections {
	return router.routes
}
