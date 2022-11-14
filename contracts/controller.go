package contracts

import "github.com/wolesgo/core/controller"

type ControllerKernelContract interface {
	GetControllers() *controller.ControllerCollections
}
