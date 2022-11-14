package kernel

import (
	"github.com/wolesgo/woles/container"
	"github.com/wolesgo/woles/contracts"
	foundation "github.com/wolesgo/woles/foundation/console"
)

type Console struct {
	foundation.Console
}

func NewConsole(app container.Container) contracts.Singleton {
	return Console{
		foundation.NewConsole(app).(foundation.Console),
	}
}
