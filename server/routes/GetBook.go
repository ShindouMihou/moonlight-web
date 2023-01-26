package routes

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"server/books"
	"server/responses"
)

func GetBook(ctx iris.Context) {
	id := ctx.Params().Get("id")
	book, err := books.Find(id)
	if err != nil {
		if err == primitive.ErrInvalidHex {
			ctx.StatusCode(iris.StatusBadRequest)
			_ = ctx.JSON(responses.InvalidPayload)
			return
		}
		ctx.StatusCode(iris.StatusInternalServerError)
		golog.Error("encountered an error while find book : ", err)
		return
	}
	if book == nil {
		ctx.StatusCode(iris.StatusNotFound)
		return
	}
	_ = ctx.JSON(book)
}
