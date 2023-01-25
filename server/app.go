package main

import (
	"github.com/kataras/golog"
	"server/metadata"
	"server/modules"
	"server/modules/server"
	"server/tokens"
)

func main() {
	golog.Info("starting ", metadata.AppName, " v", metadata.Version)
	modules.InitEnv()
	tokens.Test()
	modules.InitMongo()
	server.InitIris()
}
