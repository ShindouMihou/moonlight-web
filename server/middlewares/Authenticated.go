package middlewares

import (
	"github.com/kataras/iris/v12"
	"server/modules/regex"
	"server/responses"
	"server/tokens"
)

func Authenticated(ctx iris.Context) {
	headerValue := ctx.GetHeader("Authorization")
	if headerValue == "" {
		ctx.StatusCode(iris.StatusUnauthorized)
		_ = ctx.JSON(responses.Unauthorized)
		return
	}
	authorization := regex.BearerPrefixRegex.ReplaceAllLiteralString(headerValue, "")
	valid := tokens.Validate(authorization)
	if !valid {
		ctx.StatusCode(iris.StatusUnauthorized)
		_ = ctx.JSON(responses.Unauthorized)
		return
	}
	ctx.Next()
}
