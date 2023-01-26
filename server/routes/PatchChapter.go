package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/mongo"
	"server/chapters"
	"server/responses"
)

type iPatchChapterRequest struct {
	Title    string `json:"title" validate:"max=128"`
	Contents string `json:"contents"`
}

func PatchChapter(ctx iris.Context) {
	id := ctx.Params().GetStringTrim("id")
	var request iPatchChapterRequest
	err := ctx.ReadJSON(&request)
	if err != nil {
		if _, ok := err.(*validator.ValidationErrors); ok {
			ctx.StatusCode(iris.StatusInternalServerError)
			golog.Error("encountered an error while validating patch chapter request : ", err)
			return
		}
		ctx.StatusCode(iris.StatusBadRequest)
		_ = ctx.JSON(responses.InvalidPayload)
		return
	}

	err = chapters.Update(id, request.Title, request.Contents)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.StatusCode(iris.StatusNotFound)
			return
		}
		if err == chapters.ErrNoUpdateContents {
			ctx.StatusCode(iris.StatusBadRequest)
			_ = ctx.JSON(responses.InvalidPayload)
			return
		}
		ctx.StatusCode(iris.StatusInternalServerError)
		golog.Error("encountered an error while updating chapter : ", err)
		return
	}
	ctx.StatusCode(iris.StatusNoContent)
}
