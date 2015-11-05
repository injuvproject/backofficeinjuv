package security

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/injuvproject/backofficeinjuv/config"
)

func Encode(claims map[string]interface{}) string {

	encoder := jwt.New(jwt.SigningMethodHS256)
	encoder.Claims = claims
	tokenString, err := encoder.SignedString([]byte(config.SecretPassword))

	if err != nil {
		panic(err)
	}

	return tokenString
}

func Decode(tokenString string) (map[string]interface{}, bool) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretPassword), nil
	})

	return token.Claims, err == nil && token.Valid
}
