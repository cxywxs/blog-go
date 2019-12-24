package router

import "github.com/kataras/iris/v12"

//
type Route struct {
	Name   string
	routes []Route
}

func LoginParty(app iris.Application) {
	app.Party("")
}
