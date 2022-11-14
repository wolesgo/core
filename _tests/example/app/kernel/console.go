package kernel

import (
	"github.com/wolesgo/core/container"
	"github.com/wolesgo/core/contracts"
	foundation "github.com/wolesgo/core/foundation/console"
)

type Console struct {
	foundation.Console
}

func NewConsole(app container.Container) contracts.Singleton {
	return Console{
		foundation.NewConsole(app).(foundation.Console),
	}
}
