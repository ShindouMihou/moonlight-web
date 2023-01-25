package app

import (
	"github.com/kataras/iris/v12/context"
	"server/method"
	"server/routes"
)

type Route struct {
	Method      string
	Path        string
	Handler     context.Handler
	Middlewares []context.Handler
	Routes      []Route
}

var Routes = []Route{
	{Method: method.GetRoute, Path: "/", Handler: routes.GetIndex},
}
