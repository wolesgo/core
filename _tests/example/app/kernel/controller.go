package kernel

import (
	"github.com/wolesgo/core/container"
	"github.com/wolesgo/core/contracts"
	"github.com/wolesgo/core/controller"

	foundation "github.com/wolesgo/core/foundation/controller"

	"github.com/wolesgo/core/_tests/example"
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
