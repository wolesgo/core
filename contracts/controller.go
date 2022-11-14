package contracts

import "github.com/wolesgo/woles/controller"

type ControllerKernelContract interface {
	GetControllers() *controller.ControllerCollections
}
