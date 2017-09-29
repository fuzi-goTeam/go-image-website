package main

import (
	"flag"

	"github.com/fuzi-goTeam/go-image-website/src"
	"github.com/teambition/gear"
	"github.com/teambition/gear/logging"
)

var (
	httpAddr = flag.String("http", ":3001", "Http Addr")
)

func main() {
	// Create application
	app := gear.New()

	// app.Set("Renderer", src.Template{})
	app.Set(gear.SetRenderer, &src.Template{})
	app.UseHandler(logging.Default())

	router := src.NewRouter()
	app.UseHandler(router)
	logging.Notice("Start app listen on: " + *httpAddr)
	app.Error(app.Listen(*httpAddr))
}
