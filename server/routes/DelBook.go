package routes

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"server/books"
	"server/chapters"
	"server/responses"
)

func DelBook(ctx iris.Context) {
	book, err := books.Find(ctx.Params().GetStringTrim("id"))
	if err != nil {
		if err == primitive.ErrInvalidHex {
			ctx.StatusCode(iris.StatusBadRequest)
			_ = ctx.JSON(responses.InvalidPayload)
			return
		}
		ctx.StatusCode(iris.StatusInternalServerError)
		golog.Error("encountered an error while validating delete book : ", err)
		return
	}
	if book == nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_ = ctx.JSON(responses.InvalidPayload)
		return
	}
	err = books.Delete(*book)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		golog.Error("encountered an error while deleting book : ", err)
		return
	}
	err = chapters.DeleteAssociated(*book)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		golog.Error("encountered an error while deleting book chapters : ", err)
		return
	}
}
