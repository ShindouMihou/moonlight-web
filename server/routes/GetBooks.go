package routes

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"server/books"
	"server/responses"
)

type iGetBooksResponse struct {
	Books  []books.Book `json:"books"`
	Length int          `json:"length"`
}

func GetBooks(ctx iris.Context) {
	booksArr, err := books.All()
	if err != nil {
		if err == primitive.ErrInvalidHex {
			ctx.StatusCode(iris.StatusBadRequest)
			_ = ctx.JSON(responses.InvalidPayload)
			return
		}
		ctx.StatusCode(iris.StatusInternalServerError)
		golog.Error("encountered an error while find chapters : ", err)
		return
	}
	_ = ctx.JSON(iGetBooksResponse{
		Books:  booksArr,
		Length: len(booksArr),
	})
}
