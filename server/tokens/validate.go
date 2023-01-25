package tokens

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"server/modules"
)

func Validate(token string) bool {
	res, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(modules.EnsureEnv("MOONLIGHT_AUTH_KEY")), nil
	})
	if err != nil {
		return false
	}
	if claims, ok := res.Claims.(jwt.MapClaims); ok && res.Valid {
		err := claims.Valid()
		if err != nil {
			return false
		}
		return claims.VerifyIssuer(modules.EnsureEnv("MOONLIGHT_AUTH_ISSUER"), true)
	}
	return false
}
