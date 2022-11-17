package foundation

import (
	"github.com/wolesgo/woles/container"
	"github.com/wolesgo/woles/contracts"
	"github.com/wolesgo/woles/controller"
)

type Controller struct {
	app container.Container

	controllers *controller.Collections
}

func NewController(app container.Container, opt controller.ControllerOption) contracts.ControllerKernelContract {
	if opt.BaseModulePath != "" {
		opt.BaseModulePath = app.GetBaseModulePath(opt.BaseModulePath)
	}

	return Controller{
		app:         app,
		controllers: controller.NewCollections(opt),
	}
}

func (controller Controller) GetApplication() container.Container {
	return controller.app
}

func (controller Controller) GetControllers() *controller.Collections {
	return controller.controllers
}

func (controller Controller) bootstrap() {
}

func (controller Controller) Register() {
}
