package main

import (
	"github.com/godzillaframework/godzilla"
)

type user struct {
	User     string
	Password string
}

type aboutapi struct {
	About    string
	Github   string
	HomePage string
	Author   string
}

var usertable = []user{
	{User: "SomeOne", Password: "PasswordOne"},
	{User: "SomeOne", Password: "PasswordTwo"},
	{User: "SomeOne", Password: "PasswordTwo"},
	{User: "SomeOne", Password: "PasswordTwo"},
	{User: "SomeOne", Password: "PasswordTwo"},
	{User: "SomeOne", Password: "PasswordTwo"},
}

var infoapi = []aboutapi{
	{About: "A Simple Api Built Using Godzilla Web Framework", Github: "https://github.com/krishpranav/godzila-restapi", HomePage: "https://github.com/godzillaframework/godzilla", Author: "KrishPranav"},
}

func main() {
	gz := godzilla.New()

	/* handlers */
	gz.Get("/index", hello)
	gz.Get("/users", userhandler)
	gz.Get("/aboutapi", aboutapihandler)

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

func aboutapihandler(godz godzilla.Context) {
	godz.SendJSON(infoapi)
}
