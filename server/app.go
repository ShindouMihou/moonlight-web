package main

import (
	"github.com/kataras/golog"
	"server/metadata"
	"server/modules"
	"server/modules/server"
)

func main() {
	golog.Info("starting ", metadata.AppName, " ", metadata.Version)
	modules.InitEnv()
	modules.InitMongo()
	server.InitIris()
}
