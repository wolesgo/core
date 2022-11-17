package example

import (
	"github.com/wolesgo/woles/_tests/example/app/controller/authentication"
	"github.com/wolesgo/woles/_tests/example/app/controller/home"
	"github.com/wolesgo/woles/controller"
)

func Controllers(collections *controller.Collections) {
	collections.Add(home.NewHomeController)
	collections.Add(authentication.NewLoginController)
	collections.Add(authentication.NewLogoutController)
}
