package contracts

import router "github.com/wolesgo/woles/router/console"

type RouterKernelContract interface {
	GetConsoleRouter() *router.Router
}
