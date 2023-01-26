package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"server/books"
	"server/responses"
)

type iPutBookRequest struct {
	Name string `json:"name" validate:"required,max=128"`
}

func PutBook(ctx iris.Context) {
	var request iPutBookRequest
	err := ctx.ReadJSON(&request)
	if err != nil {
		if _, ok := err.(*validator.ValidationErrors); ok {
			ctx.StatusCode(iris.StatusInternalServerError)
			golog.Error("encountered an error while validating new book request : ", err)
			return
		}
		ctx.StatusCode(iris.StatusBadRequest)
		_ = ctx.JSON(responses.InvalidPayload)
		return
	}
	book, err := books.Insert(books.Book{Name: request.Name})
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		golog.Error("encountered an error while inserting new book request : ", err)
		return
	}
	_ = ctx.JSON(book)
}
