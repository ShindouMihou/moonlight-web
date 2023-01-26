package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"server/books"
	"server/chapters"
	"server/responses"
)

type iPutChapterRequest struct {
	Title    string `json:"title" validate:"required,max=128"`
	Contents string `json:"contents" validate:"required"`
}

func PutChapter(ctx iris.Context) {
	var request iPutChapterRequest
	err := ctx.ReadJSON(&request)
	if err != nil {
		if _, ok := err.(*validator.ValidationErrors); ok {
			ctx.StatusCode(iris.StatusInternalServerError)
			golog.Error("encountered an error while validating new chapter request : ", err)
			return
		}
		ctx.StatusCode(iris.StatusBadRequest)
		_ = ctx.JSON(responses.InvalidPayload)
		return
	}
	book, err := books.Find(ctx.Params().GetStringTrim("id"))
	if err != nil {
		if err == primitive.ErrInvalidHex {
			ctx.StatusCode(iris.StatusBadRequest)
			_ = ctx.JSON(responses.InvalidPayload)
			return
		}
		ctx.StatusCode(iris.StatusInternalServerError)
		golog.Error("encountered an error while validating find chapter book : ", err)
		return
	}
	if book == nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_ = ctx.JSON(responses.InvalidPayload)
		return
	}
	chapter, err := chapters.Insert(chapters.Chapter{
		Title:    request.Title,
		Contents: request.Contents,
		Book:     book.Id,
	})
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		golog.Error("encountered an error while insert new chapter request : ", err)
		return
	}
	_ = ctx.JSON(chapter)
}
