package kernel

import (
	"github.com/wolesgo/core/container"
	"github.com/wolesgo/core/contracts"
	foundation "github.com/wolesgo/core/foundation/router"

	"github.com/wolesgo/core/_tests/example/routing"
)

type Router struct {
	foundation.Router
}

func NewRouter(app container.Container) contracts.Singleton {
	return Router{
		foundation.NewRouter(app).(foundation.Router),
	}
}

func (router Router) Register() {
	routing.ConsoleRouting(router.Router.GetConsoleRouter())

	router.Router.Register()
}
