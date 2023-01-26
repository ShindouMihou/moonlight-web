package routes

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"server/books"
	"server/chapters"
	"server/responses"
)

type iGetChaptersResponse struct {
	Chapters []chapters.Chapter `json:"chapters"`
	Next     string             `json:"next,omitempty"`
	Length   int                `json:"length"`
}

func GetChapters(ctx iris.Context) {
	id := ctx.Params().GetStringTrim("id")
	book, err := books.Find(ctx.Params().GetStringTrim("id"))
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
	next := ctx.URLParam("next")
	chaptersArr, err := chapters.Paginate(id, next)
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
	chaptersLen := len(chaptersArr)
	var nextValue string
	if chaptersLen != 0 {
		nextValue = chaptersArr[chaptersLen-1].Id
	}

	_ = ctx.JSON(iGetChaptersResponse{
		Chapters: chaptersArr,
		Next:     nextValue,
		Length:   chaptersLen,
	})
}
