package main

import "github.com/maxence-charriere/go-app/v10/pkg/app"

func getBody() app.HTMLBody {
	return app.Body().Styles(map[string]string{
		"height":           "100%",
		"margin":           "0",
		"display":          "flex",
		"align-items":      "center",
		"justify-content":  "center",
		"background-color": "#f8f9fa",
	})
}

func getHtml() app.HTMLHtml {
	return app.Html().Styles(map[string]string{
		"height":           "100%",
		"margin":           "0",
		"display":          "flex",
		"align-items":      "center",
		"justify-content":  "center",
		"background-color": "#f8f9fa",
	})
}
