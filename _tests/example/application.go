package example

import (
	"sync"

	"github.com/wolesgo/core/container"
)

var once sync.Once

var instance *Application

type Application struct {
	container.Container

	booted bool
}

func New() *Application {
	once.Do(func() {
		instance = &Application{
			container.New(),
			false,
		}
	})
	return instance
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
