package router

import (
	"fmt"

	"github.com/wolesgo/woles"
)

type RouteController struct {
	name   string
	method string
}

type Route struct {
	router *Router

	command     string
	description string
	options     []Option

	action     func(ctx woles.Ctx)
	controller RouteController
}

func NewRoute(command string, description string) *Route {
	return &Route{
		command:     command,
		description: description,
	}
}

func (route *Route) Argument(name string, description string) *Route {
	argument := NewArgument(ArgumentConfig{}, name, description)

	fmt.Println(argument)

	return route
}

func (route *Route) Option(flags string, description string) *Route {
	option := NewOption(OptionConfig{}, flags, description)

	route.options = append(route.options, option)

	return route
}

func (route *Route) OptionWithConfig(flags string, description string, config OptionConfig) *Route {
	option := NewOption(config, flags, description)

	route.options = append(route.options, option)

	return route
}

func (route *Route) Controller(controllerName string, handle string) *Route {
	if route.action != nil {
		panic("Action cannot implemented with controller")
	}
	route.controller = RouteController{controllerName, handle}

	return route
}

func (route *Route) Group(group func(router *Router)) *Route {
	newRouter := New()

	group(newRouter)

	route.router = newRouter

	return route
}

func (route *Route) Action(action func(ctx woles.Ctx)) *Route {
	if (route.controller != RouteController{}) {
		panic("Cannot implementing with controller")
	}

	route.action = action

	return route
}

func (route Route) Handle(opts []string) {
	ctx := woles.Ctx{}

	optionValue(opts)

	// fmt.Println(route.options)

	if (route.controller == RouteController{}) {
		route.action(ctx)
	}
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

func optionValue(opts []string) {
	fmt.Println(opts)
}

func chunkSlice(slice []string, chunkSize int) [][]string {
	var chunks [][]string

	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}
