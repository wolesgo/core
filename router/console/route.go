package router

type RouteController struct {
	name   string
	method string
}

type Route struct {
	router Router

	signature   string
	description string

	controller RouteController
}

func NewRoute(signature string, description string, router Router) *Route {
	return &Route{
		router:      router,
		signature:   signature,
		description: description,
	}
}

func (route *Route) Controller(controllerName string, handle string) *Route {
	route.controller = RouteController{controllerName, handle}

	// controller := route.router.controllerCallBack(controllerName)

	// if controller.IsValid() {

	// 	selectedController := controller.MethodByName(handle)

	// 	in := []reflect.Value{reflect.ValueOf(route)}

	// 	val := selectedController.Call(in)

	// 	fmt.Println(val)
	// }

	return route
}

func (route Route) GetDescription() string {
	return route.description
}

func (route Route) GetSignature() string {
	return route.signature
}
