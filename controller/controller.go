package controller

import (
	"fmt"

	woles "github.com/wolesgo/woles"
)

type Params map[string]any

type BaseController struct {
}

func (controller BaseController) OnCreated(c woles.Ctx) {
	fmt.Println("BaseController", "OnCreated")
}
