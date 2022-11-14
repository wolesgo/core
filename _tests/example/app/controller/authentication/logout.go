package authentication

import (
	woles "github.com/wolesgo/woles"
	"github.com/wolesgo/woles/controller"
)

type LogoutController struct {
	controller.BaseController
}

func NewLogoutController() any {
	return LogoutController{
		controller.BaseController{},
	}
}

func (controller LogoutController) OnCreated(c woles.Ctx) {
	controller.BaseController.OnCreated(c)
}

func (controller LogoutController) Action(c woles.Ctx) {
}
