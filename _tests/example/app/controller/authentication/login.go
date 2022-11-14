package authentication

import (
	"fmt"

	woles "github.com/wolesgo/woles"
	"github.com/wolesgo/woles/controller"
)

type LoginController struct {
	controller.BaseController
}

func NewLoginController() any {
	return LoginController{
		controller.BaseController{},
	}
}

func (controller LoginController) OnCreated(c woles.Ctx) {
	controller.BaseController.OnCreated(c)
}

func (controller LoginController) Action(c woles.Ctx) {
	fmt.Println("LoginController@Action")
}
