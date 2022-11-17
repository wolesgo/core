package controller

import (
	"reflect"
	"regexp"
	"strings"

	woles "github.com/wolesgo/woles"
)

type Collections struct {
	baseModulePath string
	controllers    map[string]reflect.Value
}

func NewCollections(opt ControllerOption) *Collections {
	controllers := &Collections{
		controllers: make(map[string]reflect.Value),
	}

	if opt.BaseModulePath != "" {
		controllers.baseModulePath = opt.BaseModulePath
	}

	return controllers
}

func (collection Collections) Add(controllerClosure func() interface{}) reflect.Value {
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

func (collection Collections) Get(name string) reflect.Value {
	return collection.controllers[name]
}

func (controllers Collections) build(controllerClosure func() interface{}) reflect.Value {
	controller := controllerClosure()

	return reflect.ValueOf(controller)
}

func (controllers Collections) hookOnCreated(controller reflect.Value) {
	onCreatedMethod := controller.MethodByName("OnCreated")

	ctx := woles.Ctx{}

	if onCreatedMethod.IsValid() {
		onCreatedMethod.Call([]reflect.Value{reflect.ValueOf(ctx)})
	}
}
