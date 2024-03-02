package middlewares

import (
	"mdware/config"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(hp int) (string, error) {
	var data = jwt.MapClaims{}

	data["hp"] = hp
	data["iat"] = time.Now().Unix()
	data["exp"] = time.Now().Add(time.Hour * 3).Unix()

	var processToken = jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	result, err := processToken.SignedString([]byte(config.JWTSECRET))

	if err != nil {
		return "", err
	}
	return result, nil
}
