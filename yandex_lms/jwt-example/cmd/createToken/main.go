package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	const hmacSampleSecret = "super_secret_signature"
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": "user_name",
		"nbf":  now.Add(time.Minute).Unix(),
		"exp":  now.Add(5 * time.Minute).Unix(),
		"iat":  now.Unix(),
	})

	tokenString, err := token.SignedString([]byte(hmacSampleSecret))
	if err != nil {
		panic(err)
	}

	fmt.Println(tokenString)
}
