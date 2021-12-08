package main

import (
	"github.com/godzillaframework/godzilla"
)

func main() {
	gz := godzilla.New()

	gz.Get("/index", hello)
	gz.Static("/main", "./static/index.html")

	gz.Start(":8080")

}

func hello(godz godzilla.Context) {
	godz.SendString("Hello")
}
