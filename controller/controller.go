package controller

import (
	"fmt"

	"github.com/wolesgo/core/ctx"
)

type Params map[string]any

type BaseController struct {
}

func (controller BaseController) OnCreated(c ctx.Ctx) {
	fmt.Println("BaseController", "OnCreated")
}
