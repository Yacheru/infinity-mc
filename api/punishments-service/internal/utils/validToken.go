package utils

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"

	"punishments-service/internal/entities"
)

var (
	TokenMalformed = errors.New("token is malformed")
	TokenExpired   = errors.New("token is expired")
	TokenInvalid   = errors.New("token is invalid")
)

func ValidToken(t, salt string) (*entities.Claims, error) {
	claims := new(entities.Claims)

	_, err := jwt.ParseWithClaims(t, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("unexpected signing method: %v", token.Header["alg"]))
		}
		return []byte(salt), nil
	})

	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, TokenExpired
			} else {
				return nil, TokenInvalid
			}
		}
	}

	return claims, nil
}