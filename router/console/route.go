package router

type RouteController struct {
	name   string
	method string
}

type Route struct {
	router *Router

	command     string
	description string

	controller RouteController
}

func NewRoute(command string, description string) *Route {
	return &Route{
		command:     command,
		description: description,
	}
}

func (route *Route) Controller(controllerName string, handle string) *Route {
	route.controller = RouteController{controllerName, handle}

	return route
}

func (route *Route) Group(group func(router *Router)) *Route {
	newRouter := New()

	group(newRouter)

	route.router = newRouter

	return route
}

func (route Route) GetDescription() string {
	return route.description
}

func (route Route) GetCommand() string {
	return route.command
}

func (route Route) GetRouter() *Router {
	return route.router
}
