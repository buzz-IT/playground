package main

import "github.com/maxence-charriere/go-app/v10/pkg/app"

func getFooter() app.UI {
	return app.Footer().Class("mt-5 mb-3 text-muted text-center").Body(
		app.Text("© Buzz-IT GmbH"),
	)
}
