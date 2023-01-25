package router

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/core/router"
	"server/app"
	"server/method"
)

func Attach(Iris *iris.Application) {
	var mod = func(route app.Route, r *router.Route) {
		for _, middleware := range route.Middlewares {
			r.Use(middleware)
		}
	}

	var AddParty func(route app.Route, party iris.Party)
	AddParty = func(route app.Route, party iris.Party) {
		for _, additionalRoute := range route.Routes {
			// IMPORTANT: Since we cannot make use of party to add the middlewares, we'll copy them one-by-one instead.
			var middlewares = additionalRoute.Middlewares
			if route.Middlewares != nil && additionalRoute.Middlewares != nil {
				middlewares = make([]context.Handler, (len(route.Middlewares)+len(additionalRoute.Middlewares))-1)
				copy(middlewares, route.Middlewares)
				middlewares = append(middlewares, additionalRoute.Middlewares...)
			}
			if route.Middlewares != nil && additionalRoute.Middlewares == nil {
				middlewares = route.Middlewares
			}
			additionalRoute.Middlewares = middlewares

			golog.Info("adding ", additionalRoute.Method, " ", additionalRoute.Path, " route to party ", party.GetRelPath(), " with ", len(additionalRoute.Middlewares), " middlewares")
			switch additionalRoute.Method {
			case method.GetRoute:
				mod(additionalRoute, party.Get(additionalRoute.Path, additionalRoute.Handler))
			case method.HeadRoute:
				mod(additionalRoute, party.Head(additionalRoute.Path, additionalRoute.Handler))
			case method.PostRoute:
				mod(additionalRoute, party.Post(additionalRoute.Path, additionalRoute.Handler))
			case method.PutRoute:
				mod(additionalRoute, party.Put(additionalRoute.Path, additionalRoute.Handler))
			case method.DeleteRoute:
				mod(additionalRoute, party.Delete(additionalRoute.Path, additionalRoute.Handler))
			case method.PartyRoute:
				additionalParty := party.Party(additionalRoute.Path)
				AddParty(additionalRoute, additionalParty)
			}
		}
	}

	for _, route := range app.Routes {
		golog.Info("adding ", route.Method, " ", route.Path, " route with ", len(route.Middlewares), " middlewares")
		switch route.Method {
		case method.GetRoute:
			mod(route, Iris.Get(route.Path, route.Handler))
		case method.HeadRoute:
			mod(route, Iris.Head(route.Path, route.Handler))
		case method.PostRoute:
			mod(route, Iris.Post(route.Path, route.Handler))
		case method.PutRoute:
			mod(route, Iris.Put(route.Path, route.Handler))
		case method.DeleteRoute:
			mod(route, Iris.Delete(route.Path, route.Handler))
		case method.PartyRoute:
			party := Iris.Party(route.Path)
			AddParty(route, party)
		}
	}
}
