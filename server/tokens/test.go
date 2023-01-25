package tokens

import (
	"github.com/kataras/golog"
)

func Test() {
	for i := 0; i < 10; i++ {
		token, err := Create()
		if err != nil {
			golog.Fatal("failed to create token during test at index ", i, " with error: ", err)
		}
		valid := Validate(*token)
		if !valid {
			golog.Fatal("invalid token ", *token, " despite the token being newly generated.")
		}
	}
	golog.Info("token creation and validation passed")
}
