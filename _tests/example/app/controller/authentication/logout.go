package authentication

import (
	"github.com/wolesgo/core/controller"
	"github.com/wolesgo/core/ctx"
)

type LogoutController struct {
	controller.BaseController
}

func NewLogoutController() any {
	return LogoutController{
		controller.BaseController{},
	}
}

func (controller LogoutController) OnCreated(c ctx.Ctx) {
	controller.BaseController.OnCreated(c)
}

func (controller LogoutController) Action(c ctx.Ctx) {
}
