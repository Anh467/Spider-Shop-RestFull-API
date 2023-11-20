package common

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	UserID   int    `json:"userid"`
	Account  string `json:"account"`
	Password string `json:"password"`
	jwt.StandardClaims
}
