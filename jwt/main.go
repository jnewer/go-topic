package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

var MySecret = []byte("secret")

func GenToken(username string, password string) (string, error) {
	c := MyClaims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "kong",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(MySecret)
}

func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return MySecret, nil
		})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
func main() {
	//s, err := GenToken("ghz", "123")
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Printf("s: %v", s)

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImdoeiIsInBhc3N3b3JkIjoiMTIzIiwiZXhwIjoxNjU5OTY5MzY0LCJpc3MiOiJrb25nIn0.POwgYydze6Qyo4MIhvKDj_3oAVapveRJrqQbv8U5abU"

	mc, err := ParseToken(token)
	if err != nil {
		panic(err)
	}
	fmt.Printf("mc.Password: %v\n", mc.Password)
	fmt.Printf("mc.Username: %v\n", mc.Username)
}
