package models

import "github.com/golang-jwt/jwt"

type AppClaims struct { // no usamos herencia, usamos composición sobre herencia
	UserId string `json:"userId"`
	jwt.StandardClaims
}

/*
jwt.StandardClaims se compone de:
type StandardClaims struct {
    Audience  string `json:"aud,omitempty"`
    ExpiresAt int64  `json:"exp,omitempty"`
    Id        string `json:"jti,omitempty"`
    IssuedAt  int64  `json:"iat,omitempty"`
    Issuer    string `json:"iss,omitempty"`
    NotBefore int64  `json:"nbf,omitempty"`
    Subject   string `json:"sub,omitempty"`
}

al posicionar mouse sobre jwt.StandardClaims muestra la struct
*/
