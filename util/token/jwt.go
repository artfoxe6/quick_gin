package token

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"quick_gin/config/env"
)

// 生成token
func CreateJwtToken(data map[string]interface{}) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, (jwt.MapClaims)(data))
	tokenString, err := t.SignedString([]byte(env.Jwt.Key))
	if err != nil {
		log.Fatalf("%v", err)
	}
	return tokenString, nil
}

// 验证token
func VerifyJwtToken(tokenString string) error {
	t, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(env.Jwt.Key), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		return claims.Valid()
	} else {
		return err
	}
}
