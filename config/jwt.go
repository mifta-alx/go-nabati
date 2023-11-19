package config

import "github.com/golang-jwt/jwt/v5"

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
