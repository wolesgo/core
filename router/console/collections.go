package router

import "errors"

type RouteCollections map[string]*Route

func NewRouteCollections() *RouteCollections {
	newRouteCollections := make(RouteCollections)
	return &newRouteCollections
}

func (c *RouteCollections) Add(route *Route) *Route {
	(*c)[route.GetCommand()] = route

	return route
}

func (c RouteCollections) RouteByName(name string) (*Route, error) {
	if c[name] != nil {
		return c[name], nil
	}
	return c[name], errors.New("No Route Found.")
}

func (c RouteCollections) Get(route interface{}) Route {
	return route.(Route)
}

func (c RouteCollections) GetRoutes() RouteCollections {
	return c
}
