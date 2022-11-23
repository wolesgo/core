package foundation

import (
	"github.com/wolesgo/woles/container"
	"github.com/wolesgo/woles/contracts"
)

type Console struct {
	app container.Container
}

func NewConsole(app container.Container) contracts.ConsoleKernelContract {
	return Console{
		app,
	}
}

func (console Console) Handle(args []string) {
	console.bootstrap()

	routerKenel := console.GetApplication().Resolve((*contracts.RouterKernelContract)(nil))

	kernel := routerKenel.(contracts.Singleton)

	router := kernel.(contracts.RouterKernelContract).GetConsoleRouter()

	router.MatchRoute(args)
	// console.HandleRoute(router)
}

func (console Console) GetApplication() *container.Container {
	return &console.app
}

func (console Console) bootstrap() {
}

func (console Console) Register() {
}
