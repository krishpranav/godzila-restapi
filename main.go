package main

import (
	"github.com/godzillaframework/godzilla"
)

func main() {
	gz := godzilla.New()

	gz.Get("/index", hello)

	gz.Start(":8080")
}

func hello(ctx godzilla.Context) {
	ctx.SendString("Hello")
}
