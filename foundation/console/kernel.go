package foundation

import (
	"github.com/wolesgo/core/container"
	"github.com/wolesgo/core/contracts"
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
}

func (console Console) GetApplication() container.Container {
	return console.app
}

func (console Console) bootstrap() {
}

func (console Console) Register() {
}
