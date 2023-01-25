package routes

import (
	"github.com/kataras/iris/v12"
	"server/metadata"
	"time"
)

type iGetIndexResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Version   float64   `json:"version"`
}

func GetIndex(ctx iris.Context) {
	_ = ctx.JSON(iGetIndexResponse{
		Timestamp: time.Now(),
		Version:   metadata.Version,
	})
}
