package home

import (
	"fmt"

	woles "github.com/wolesgo/woles"
	"github.com/wolesgo/woles/controller"
)

type HomeController struct {
	controller.BaseController
}

func NewHomeController() any {
	return HomeController{
		BaseController: controller.BaseController{},
	}
}

func (controller HomeController) OnCreated(c woles.Ctx) {
	controller.BaseController.OnCreated(c)
}

func (controller HomeController) Action(c woles.Ctx) {
	fmt.Println("HomeController@Action", c)
}
