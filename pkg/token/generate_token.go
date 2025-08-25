package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(data *UserAuthToken) (string, error) {
	token := jwt.New(jwt.SigningMethodRS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["data"] = data
	claims["iss"] = "https://www.kemenkeu.go.id"
	claims["iat"] = time.Now()
	if jwtConfig == nil {
		panic("jwtConfig is nil, please initialize JWT config before use")
	}
	claims["exp"] = time.Now().Add(time.Duration(jwtConfig.jwtLifeTime) * time.Second).Unix()

	return token.SignedString(jwtConfig.privateKey)
}

func GenerateRefreshToken(id uint64) (string, error) {
	token := jwt.New(jwt.SigningMethodRS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["data"] = map[string]uint64{
		"id": id,
	}
	claims["iss"] = "https://www.kemenkeu.go.id"
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Duration(jwtConfig.jwtLifeTime) * time.Second).Unix()

	return token.SignedString(jwtConfig.privateKey)
}
