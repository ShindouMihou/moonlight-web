package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"server/router"
)

var Iris *iris.Application

func InitIris() {
	config := iris.WithConfiguration(iris.Configuration{
		DisableStartupLog: true,
	})

	Iris = iris.New()
	Iris.Validator = validator.New()
	Iris.Use(logger.New(logger.DefaultConfig()))
	router.Attach(Iris)

	golog.Info("launching iris at http://0.0.0.0:7779")
	err := Iris.Listen(":7779", config)
	if err != nil {
		golog.Fatal("an error occurred while trying to start iris: ", err)
	}
}
