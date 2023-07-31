package middleware

import (
	"errors"
	"project/constant"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userId int, uuid string, name string) (string, int64, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["uuid"] = uuid
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(constant.SECRET_JWT))
	if err != nil {
		return "", 0, err
	}

	expirationTime, ok := claims["exp"].(int64)
	if !ok {
		return "", 0, errors.New("failed to get expiration time")
	}

	return tokenString, expirationTime, nil
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(constant.SECRET_JWT), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to get claims")
	}

	return claims, nil
}
