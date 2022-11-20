package models

import "github.com/golang-jwt/jwt"

// https://github.com/dgrijalva/jwt-go THIS REPOSITORY IS NO LONGER MAINTANED
type AppClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}
