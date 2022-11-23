package router

import (
	"errors"
	"reflect"
)

type opts = []string

type Router struct {
	routes *RouteCollections

	controllerCallBack func(controllerName string) reflect.Value
}

func New() *Router {
	return &Router{
		routes: NewRouteCollections(),
	}
}

func (router *Router) RegisterControllerCallBack(callback func(controllerName string) reflect.Value) {
	router.controllerCallBack = callback
}

func (router *Router) Command(command string, description string) *Route {
	_, err := router.routes.RouteByName(command)
	if err == nil {
		panic(errors.New("Command already exists."))
	}

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
	return router.routes.GetRoutes()
}

func (router Router) MatchRoute(args []string) {
	route, opts, err := router.HandleRoute(router, args, make(opts, 0))

	if err == nil {
		route.Handle(opts)
	}
}

func (router Router) HandleRoute(r Router, args []string, opts []string) (*Route, opts, error) {
	routes := router.GetRoutes()

	if len(args) > 0 {
		matchingRoute, err := routes.RouteByName(args[0])
		if err != nil {
			return nil, nil, errors.New("No Route Found.")
		}

		router := matchingRoute.GetRouter()

		if router != nil {
			return router.HandleRoute(*router, args[1:], args[1:])
		}

		return matchingRoute, opts[1:], nil
	}

	return nil, nil, errors.New("No Route Found.")
}
