package authentication

import (
	"fmt"

	"github.com/wolesgo/core/controller"
	"github.com/wolesgo/core/ctx"
)

type LoginController struct {
	controller.BaseController
}

func NewLoginController() any {
	return LoginController{
		controller.BaseController{},
	}
}

func (controller LoginController) OnCreated(c ctx.Ctx) {
	controller.BaseController.OnCreated(c)
}

func (controller LoginController) Action(c ctx.Ctx) {
	fmt.Println("LoginController@Action")
}
