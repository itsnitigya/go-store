package main

import (
	"github.com/itsnitigya/go-store/app"
)

func main() {
	// config := config.GetConfig()

	app := &app.App{}
	app.Initialize()
	app.Run(":3000")
}
