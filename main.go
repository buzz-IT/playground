package main

import (
	"log"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type hello struct {
	app.Compo
}

// The Render method is where the component appearance is defined. Here, a
// "Hello World!" is displayed as a heading.
func (h *hello) Render() app.UI {
	return app.H1().Text("Hello World!")
}

// The main function is the entry point where the app is configured and started.
// It is executed in 2 different environments: A client (the web browser) and a
// server.
func main() {
	// The first thing to do is to associate the hello component with a path.
	//
	// This is done by calling the Route() function,  which tells go-app what
	// component to display for a given path, on both client and server-side.
	app.Route("/", func() app.Composer { return newUiIp() })

	// Once the routes set up, the next thing to do is to either launch the app
	// or the server that serves the app.
	//
	// When executed on the client-side, the RunWhenOnBrowser() function
	// launches the app,  starting a loop that listens for app events and
	// executes client instructions. Since it is a blocking call, the code below
	// it will never be executed.
	//
	// When executed on the server-side, RunWhenOnBrowser() does nothing, which
	// lets room for server implementation without the need for precompiling
	// instructions.
	app.RunWhenOnBrowser()

	err := app.GenerateStaticWebsite(".", &app.Handler{
		Name:        "IP Toolbox",
		Description: "IP Toolbox",
		Styles: []string{
			"https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css",
			"/web/custom.css",
		},
		Body: getBody,
		HTML: getHtml,
		Icon: app.Icon{
			Default: "/web/logo.png",
			SVG:     "/web/logo.svg",
		},
		Keywords: []string{
			"IP Tool",
			"ip",
			"buzz-it",
			"buzz-it gmbh",
			"buzz-IT GmbH",
		},
		LoadingLabel: "IP Tools...",
		ThemeColor:   "#fbb800ff",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Finally, launching the server that serves the app is done by using the Go
	// standard HTTP package.
	//
	// The Handler is an HTTP handler that serves the client and all its
	// required resources to make it work into a web browser. Here it is
	// configured to handle requests with a path that starts with "/".
	// 	http.Handle("/", &app.Handler{
	// 		Name:        "IP Tool",
	// 		Description: "An Hello World! example",
	// 		Styles: []string{
	// 			"https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css",
	// 		},
	// 		Body: getBody,
	// 		HTML: getHtml,
	// 	})

	//	if err := http.ListenAndServe(":8000", nil); err != nil {
	//		log.Fatal(err)
	//	}
}
