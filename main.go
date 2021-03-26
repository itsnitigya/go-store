package main

import (
	"github.com/itsnitigya/go-store/app"
	"github.com/itsnitigya/go-store/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
