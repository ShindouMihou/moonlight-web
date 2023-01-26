package routes

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"server/modules"
	"server/responses"
	"server/tokens"
	"time"
)

var RootUsername string
var RootPassword string

type iPutTokenRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type iPutTokenResponse struct {
	Token     string    `json:"token"`
	ExpiresIn time.Time `json:"expiresIn"`
}

func PutToken(ctx iris.Context) {
	if RootUsername == "" || RootPassword == "" {
		RootUsername = modules.EnsureEnv("MOONLIGHT_ROOT_USERNAME")
		RootPassword = modules.EnsureEnv("MOONLIGHT_ROOT_PASSWORD")
	}
	var request iPutTokenRequest
	err := ctx.ReadJSON(&request)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_ = ctx.JSON(responses.InvalidPayload)
		return
	}
	if request.Username != RootUsername || request.Password != RootPassword {
		ctx.StatusCode(iris.StatusUnauthorized)
		_ = ctx.JSON(responses.InvalidAuthentication)
		return
	}
	token, expiry, err := tokens.Create()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		golog.Error("encountered an error while creating token : ", err)
		return
	}
	_ = ctx.JSON(iPutTokenResponse{Token: *token, ExpiresIn: *expiry})
	return
}
