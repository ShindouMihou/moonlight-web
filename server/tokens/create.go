package tokens

import (
	"github.com/golang-jwt/jwt/v4"
	"server/modules"
	"time"
)

func Create() (*string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:    modules.EnsureEnv("MOONLIGHT_AUTH_ISSUER"),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(modules.EnsureEnv("MOONLIGHT_AUTH_KEY")))
	if err != nil {
		return nil, err
	}
	return &signed, nil
}
