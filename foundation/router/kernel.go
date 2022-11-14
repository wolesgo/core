package foundation

import (
	"reflect"

	"github.com/wolesgo/woles/contracts"
	foundation "github.com/wolesgo/woles/foundation/controller"

	"github.com/wolesgo/woles/container"

	consoleRouter "github.com/wolesgo/woles/router/console"
)

type Router struct {
	app container.Container

	consoleRouter *consoleRouter.Router
}

func NewRouter(app container.Container) contracts.RouterKernelContract {
	return Router{
		app,

		&consoleRouter.Router{},
	}
}

func (router Router) GetApplication() *container.Container {
	return &router.app
}

func (router Router) GetConsoleRouter() *consoleRouter.Router {
	return router.consoleRouter
}

func (router Router) bootstrap() {
}

func (router Router) Register() {

	router.consoleRouter.RegisterControllerCallBack(router.RegisterControllerCallBack)
}

func (router Router) RegisterControllerCallBack(controllerName string) reflect.Value {
	var controllerKernel = router.GetApplication().Resolve((*contracts.ControllerKernelContract)(nil))

	var kernel = controllerKernel.(contracts.Singleton)

	var controllers = kernel.(foundation.Controller).GetControllers()

	var handlerController = controllers.Get(controllerName)

	return handlerController
}
