package routes

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"server/chapters"
	"server/responses"
)

func GetChapter(ctx iris.Context) {
	id := ctx.Params().Get("id")
	chapter, err := chapters.Find(id)
	if err != nil {
		if err == primitive.ErrInvalidHex {
			ctx.StatusCode(iris.StatusBadRequest)
			_ = ctx.JSON(responses.InvalidPayload)
			return
		}
		ctx.StatusCode(iris.StatusInternalServerError)
		golog.Error("encountered an error while find chapter : ", err)
		return
	}
	if chapter == nil {
		ctx.StatusCode(iris.StatusNotFound)
		return
	}
	_ = ctx.JSON(chapter)
}
