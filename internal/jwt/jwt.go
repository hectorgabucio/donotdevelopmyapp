package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtProvider interface {
	CreateToken(userId string, secret []byte, expires time.Duration) (string, error)
	DecodeToken(token string, secret []byte) (string, error)
}

type JwtImpl struct {
}

func (j JwtImpl) CreateToken(userid string, secret []byte, expires time.Duration) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(expires).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString(secret)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (j JwtImpl) DecodeToken(token string, secret []byte) (string, error) {
	claims := jwt.MapClaims{}
	jwt, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return "", err
	}

	if !jwt.Valid {
		return "", fmt.Errorf("Token not valid")
	}

	for key, val := range claims {
		if key == "user_id" {
			return val.(string), nil
		}

	}

	return "", fmt.Errorf("Could not find user id claim in decoded token")

}
