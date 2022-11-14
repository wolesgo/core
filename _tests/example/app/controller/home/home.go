package home

import (
	"fmt"

	"github.com/wolesgo/core/controller"
	"github.com/wolesgo/core/ctx"
)

type HomeController struct {
	controller.BaseController
}

func NewHomeController() any {
	return HomeController{
		BaseController: controller.BaseController{},
	}
}

func (controller HomeController) OnCreated(c ctx.Ctx) {
	controller.BaseController.OnCreated(c)
}

func (controller HomeController) Action(c ctx.Ctx) {
	fmt.Println("HomeController@Action", c)
}
