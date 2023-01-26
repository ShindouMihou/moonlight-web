package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/mongo"
	"server/books"
	"server/responses"
)

type iPatchBookRequest struct {
	Name string `json:"name" validate:"max=128"`
}

func PatchBook(ctx iris.Context) {
	id := ctx.Params().GetStringTrim("id")
	var request iPatchBookRequest
	err := ctx.ReadJSON(&request)
	if err != nil {
		if _, ok := err.(*validator.ValidationErrors); ok {
			ctx.StatusCode(iris.StatusInternalServerError)
			golog.Error("encountered an error while validating patch book request : ", err)
			return
		}
		ctx.StatusCode(iris.StatusBadRequest)
		_ = ctx.JSON(responses.InvalidPayload)
		return
	}
	err = books.Update(id, request.Name)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.StatusCode(iris.StatusNotFound)
			return
		}
		if err == books.ErrNoUpdateContents {
			ctx.StatusCode(iris.StatusBadRequest)
			_ = ctx.JSON(responses.InvalidPayload)
			return
		}
		ctx.StatusCode(iris.StatusInternalServerError)
		golog.Error("encountered an error while updating book : ", err)
		return
	}
	ctx.StatusCode(iris.StatusNoContent)
}
