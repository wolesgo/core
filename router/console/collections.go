package router

type RouteCollections []Route

func (routes *RouteCollections) Add(route Route) *Route {
	*routes = append(*routes, route)
	return &route
}

func (routes RouteCollections) Get(route interface{}) Route {
	return route.(Route)
}
