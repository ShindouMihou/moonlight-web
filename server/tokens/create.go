package tokens

import (
	"github.com/golang-jwt/jwt/v4"
	"server/modules"
	"time"
)

func Create() (*string, *time.Time, error) {
	expiry := time.Now().Add(30 * 24 * time.Hour)
	claims := jwt.RegisteredClaims{
		Issuer:    modules.EnsureEnv("MOONLIGHT_AUTH_ISSUER"),
		ExpiresAt: jwt.NewNumericDate(expiry),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(modules.EnsureEnv("MOONLIGHT_AUTH_KEY")))
	if err != nil {
		return nil, nil, err
	}
	return &signed, &expiry, nil
}
