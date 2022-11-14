package controller

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/wolesgo/core/ctx"
)

type ControllerCollections struct {
	baseModulePath string
	controllers    map[string]reflect.Value
}

func NewCollections(opt ControllerOption) *ControllerCollections {
	controllers := &ControllerCollections{
		controllers: make(map[string]reflect.Value),
	}

	if opt.BaseModulePath != "" {
		controllers.baseModulePath = opt.BaseModulePath
	}

	return controllers
}

func (collection ControllerCollections) Add(controllerClosure func() interface{}) reflect.Value {
	controller := collection.build(controllerClosure)

	pkgPath := regexp.MustCompile(`\/+$`).ReplaceAllString(
		regexp.MustCompile(`^\/+`).ReplaceAllString(
			strings.Replace(
				controller.Type().PkgPath(),
				collection.baseModulePath, "",
				-1,
			),
			"",
		),
		"",
	) + "/"

	collection.controllers[pkgPath+controller.Type().Name()] = controller

	collection.hookOnCreated(controller)

	return controller
}

func (collection ControllerCollections) Get(name string) reflect.Value {
	return collection.controllers[name]
}

func (controllers ControllerCollections) build(controllerClosure func() interface{}) reflect.Value {
	controller := controllerClosure()

	return reflect.ValueOf(controller)
}

func (controllers ControllerCollections) hookOnCreated(controller reflect.Value) {
	onCreatedMethod := controller.MethodByName("OnCreated")

	ctx := ctx.Ctx{}

	if onCreatedMethod.IsValid() {
		onCreatedMethod.Call([]reflect.Value{reflect.ValueOf(ctx)})
	}
}
