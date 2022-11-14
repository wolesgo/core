package bootstrap

import (
	"io/ioutil"
	"sync"

	"github.com/wolesgo/woles/contracts"
	"github.com/wolesgo/woles/foundation"
	"golang.org/x/mod/modfile"

	"github.com/wolesgo/woles/_tests/example/app/kernel"
)

var once sync.Once

var app *foundation.Application

func New(Args []string) *foundation.Application {
	once.Do(func() {
		app = foundation.New()

		goModBytes, err := ioutil.ReadFile("go.mod")
		if err != nil {
		}

		app.BaseModulePath(modfile.ModulePath(goModBytes))
	})

	app.Singleton(
		(*contracts.ControllerKernelContract)(nil),
		kernel.NewController,
	)

	app.Singleton(
		(*contracts.ConsoleKernelContract)(nil),
		kernel.NewConsole,
	)

	app.Singleton(
		(*contracts.RouterKernelContract)(nil),
		kernel.NewRouter,
	)

	app.Boot()

	var kernel = app.Make((*contracts.ConsoleKernelContract)(nil)).(kernel.Console)

	kernel.Handle(Args[1:])

	return app
}
