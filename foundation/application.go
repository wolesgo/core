package foundation

import (
	"github.com/wolesgo/woles/container"
)

type Application struct {
	container.Container

	booted bool
}

func New() *Application {
	return &Application{
		container.New(),
		false,
	}
}

func (application Application) Register() {
	application.Container.Register()
}

func (application Application) Boot() bool {
	if application.booted {
		return false
	}

	application.Register()

	return true
}
