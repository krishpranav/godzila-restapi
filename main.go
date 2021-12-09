package main

import (
	"github.com/godzillaframework/godzilla"
)

type user struct {
	User     string
	Password string
}

var usertable = []user{
	{User: "SomeOne", Password: "PasswordOne"},
	{User: "SomeOne", Password: "PasswordTwo"},
	{User: "SomeOne", Password: "PasswordTwo"},
	{User: "SomeOne", Password: "PasswordTwo"},
	{User: "SomeOne", Password: "PasswordTwo"},
	{User: "SomeOne", Password: "PasswordTwo"},
}

func main() {
	gz := godzilla.New()

	/* handlers */
	gz.Get("/index", hello)
	gz.Get("/users", userhandler)

	/* static web server */
	gz.Static("/main", "./static/index.html")

	gz.Start(":8080")

}

func hello(godz godzilla.Context) {
	godz.SendString("Hello")
}

func userhandler(godz godzilla.Context) {
	godz.SendJSON(usertable)
}
