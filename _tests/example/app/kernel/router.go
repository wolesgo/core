package kernel

import (
	"github.com/wolesgo/woles/container"
	"github.com/wolesgo/woles/contracts"
	foundation "github.com/wolesgo/woles/foundation/router"

	"github.com/wolesgo/woles/_tests/example/routing"
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
	routing.ConsoleRouting(router.GetConsoleRouter())

	router.Router.Register()
}
