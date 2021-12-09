package main

import (
	"fmt"

	"github.com/godzillaframework/godzilla"
)

// var usertable = []user.users{
// 	{Username: "USername one", Password: "Password"},
// 	{Username: "USername one", Password: "Password"},
// 	{Username: "USername one", Password: "Password"},
// 	{Username: "USername one", Password: "Password"},
// 	{Username: "USername one", Password: "Password"},
// }

func main() {
	// gz := godzilla.New()

	// gz.Get("/index", hello)
	// gz.Static("/main", "./static/index.html")

	// gz.Start(":8080")

	fmt.Println(usertable)

}

func hello(godz godzilla.Context) {
	godz.SendString("Hello")
}

func userhandler(godz godzilla.Context) {

}
