package app

import (
	"github.com/kataras/iris/v12/context"
	"server/method"
	"server/middlewares"
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
	{Method: method.PutRoute, Path: "/token", Handler: routes.PutToken},
	{
		Method:      method.PartyRoute,
		Path:        "/books",
		Middlewares: []context.Handler{middlewares.Authenticated},
		Routes: []Route{
			{Method: method.GetRoute, Handler: routes.GetBooks},
			{Method: method.PutRoute, Handler: routes.PutBook},
			{
				Method: method.PartyRoute,
				Path:   "/{id}",
				Routes: []Route{
					{Method: method.DeleteRoute, Handler: routes.DelBook},
					{Method: method.PatchRoute, Handler: routes.PatchBook},
					{
						Method: method.PartyRoute,
						Path:   "/chapters",
						Routes: []Route{
							{Method: method.GetRoute, Handler: routes.GetChapters},
							{Method: method.PutRoute, Handler: routes.PutChapter},
						},
					},
				},
			},
		},
	},
	{
		Method:      method.PartyRoute,
		Path:        "/chapters",
		Middlewares: []context.Handler{middlewares.Authenticated},
		Routes: []Route{
			{
				Method: method.PartyRoute,
				Path:   "/{id}",
				Routes: []Route{
					{Method: method.DeleteRoute, Handler: routes.DelChapter},
					{Method: method.GetRoute, Handler: routes.GetChapter},
					{Method: method.PatchRoute, Handler: routes.PatchChapter},
				},
			},
		},
	},
}
