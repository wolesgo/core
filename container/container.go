package container

import (
	"fmt"
	"reflect"

	"github.com/wolesgo/woles/contracts"
	"golang.org/x/mod/modfile"
)

type Container struct {
	booted bool

	baseModulePath string

	bindings []string

	bindingConcretes map[string]contracts.Singleton
}

func New() Container {
	bindingConcretes := make(map[string]contracts.Singleton)

	return Container{
		booted:           false,
		baseModulePath:   "",
		bindingConcretes: bindingConcretes,
	}
}

func (container Container) build(concrete interface{}) contracts.Singleton {
	if concrete, isInstantiable := concrete.(func(App Container) contracts.Singleton); isInstantiable {
		return concrete(container)
	}

	return nil
}

func (container *Container) bind(abstract interface{}, concrete interface{}) {
	var instantiable contracts.Singleton = container.build(concrete)

	abstractName := reflect.TypeOf(abstract).Elem().Name()

	if concrete != nil {
		container.bindings = append(container.bindings, abstractName)
		container.bindingConcretes[abstractName] = instantiable
	}
}

func (container *Container) Resolve(abstract interface{}) interface{} {
	abstractName := reflect.TypeOf(abstract).Elem().Name()

	return container.bindingConcretes[abstractName]
}

func (container *Container) Singleton(abstract interface{}, concrete interface{}) {
	container.bind(abstract, concrete)
}

func (container Container) Register() {
	for _, binding := range container.bindings {
		if concrete, isInstantiable := container.bindingConcretes[binding]; isInstantiable {
			concrete.Register()
		}
	}
}

func (container *Container) BaseModulePathByMod(goModBytes []byte, err error) string {
	if err != nil {
	}

	container.baseModulePath = modfile.ModulePath(goModBytes)

	return container.baseModulePath
}

func (container Container) GetBaseModulePath(appendModulePath string) string {
	baseModulePath := container.baseModulePath
	if appendModulePath != "" {
		baseModulePath += "/" + appendModulePath
	}
	return baseModulePath
}

func (container *Container) Make(abstract interface{}) interface{} {
	return container.Resolve(abstract)
}

func (container Container) Hello() {
	fmt.Println("--Hello--")
}
