package routes

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"server/chapters"
	"server/responses"
)

func DelChapter(ctx iris.Context) {
	chapter, err := chapters.Find(ctx.Params().GetStringTrim("id"))
	if err != nil {
		if err == primitive.ErrInvalidHex {
			ctx.StatusCode(iris.StatusBadRequest)
			_ = ctx.JSON(responses.InvalidPayload)
			return
		}
		ctx.StatusCode(iris.StatusInternalServerError)
		golog.Error("encountered an error while validating delete chapter : ", err)
		return
	}
	if chapter == nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_ = ctx.JSON(responses.InvalidPayload)
		return
	}
	err = chapters.Delete(*chapter)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		golog.Error("encountered an error while deleting book : ", err)
		return
	}
}
