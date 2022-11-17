package routing

import (
	router "github.com/wolesgo/woles/router/console"
)

func ConsoleRouting(r *router.Router) {

	r.Command(
		"image",
		"to get image",
	).Group(func(r *router.Router) {
		r.Command(
			"ls",
			"To ls.",
		).Group(func(r *router.Router) {
			r.Command(
				"all",
				"to all.",
			).Controller(
				"home/HomeController",
				"Action",
			)

			r.Command(
				"used",
				"To used.",
			).Controller(
				"home/HomeController",
				"Action",
			)
		})
	})

	r.Command(
		"login",
		"To login into the system",
	).Controller(
		"LoginController",
		"Action",
	)

	r.Command(
		"logout",
		"To logout from the system",
	).Controller(
		"LoginController",
		"Action",
	)
}
