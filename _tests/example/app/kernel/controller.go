package kernel

import (
	"github.com/wolesgo/woles/container"
	"github.com/wolesgo/woles/contracts"
	"github.com/wolesgo/woles/controller"

	foundation "github.com/wolesgo/woles/foundation/controller"

	"github.com/wolesgo/woles/_tests/example"
)

type Controller struct {
	foundation.Controller
}

func NewController(app container.Container) contracts.Singleton {
	return Controller{
		foundation.NewController(app, controller.ControllerOption{
			BaseModulePath: "_tests/example/app/controller",
		}).(foundation.Controller),
	}
}

func (kernel Controller) Register() {
	example.Controllers(kernel.GetControllers())

	kernel.Controller.Register()
}
