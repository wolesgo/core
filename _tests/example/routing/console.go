package routing

import (
	"github.com/wolesgo/woles"
	router "github.com/wolesgo/woles/router/console"
)

func ConsoleRouting(r *router.Router) {

	r.Command("image", "to get image").
		Group(func(r *router.Router) {
			r.Command("ls", "To ls.").
				Group(func(r *router.Router) {
					r.Command("all", "to all.").
						Argument("<fizz>", "output extra debugging").
						Option("-d, --debug", "output extra debugging").
						Option("-f <char>", "foo").
						Option("--bar <char>", "bar").
						Option("-s, --separator <char>", "bar").
						Action(func(ctx woles.Ctx) {
							// fmt.Println("all", ctx)
						})
						// .Controller("home/HomeController", "Action")

					r.Command("used", "To used.").
						Option("--debug", "output extra debugging").
						Controller("home/HomeController", "Action")
				})
		})

	// r.Command(
	// 	"image",
	// 	"To login into the system",
	// ).Controller(
	// 	"LoginController",
	// 	"Action",
	// )

	r.Command("login", "To login into the system").
		Controller("LoginController", "Action")

	r.Command("logout", "To logout from the system").
		Controller("LoginController", "Action")
}
