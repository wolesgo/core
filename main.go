package main

import (
	"io/ioutil"
	"os"

	"github.com/wolesgo/core/contracts"
	"golang.org/x/mod/modfile"

	"github.com/wolesgo/core/_tests/example"

	"github.com/wolesgo/core/_tests/example/app/kernel"
)

func main() {
	bootstrap := example.New()

	goModBytes, err := ioutil.ReadFile("go.mod")
	if err != nil {
	}

	bootstrap.BaseModulePath(modfile.ModulePath(goModBytes))

	bootstrap.Singleton(
		(*contracts.ControllerKernelContract)(nil),
		kernel.NewController,
	)

	bootstrap.Singleton(
		(*contracts.ConsoleKernelContract)(nil),
		kernel.NewConsole,
	)

	bootstrap.Singleton(
		(*contracts.RouterKernelContract)(nil),
		kernel.NewRouter,
	)

	bootstrap.Boot()

	var kernel = bootstrap.Make((*contracts.ConsoleKernelContract)(nil)).(kernel.Console)

	kernel.Handle(os.Args[1:])
}
