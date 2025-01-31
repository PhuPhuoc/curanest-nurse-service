package common

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type JWTx struct {
	secret string
}

func NewJWTx(secret string) *JWTx {
	return &JWTx{
		secret: secret,
	}
}

func (j *JWTx) ParseToken(ctx context.Context, tokenString string) (map[string]interface{}, error) {
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	return claims, nil
}
