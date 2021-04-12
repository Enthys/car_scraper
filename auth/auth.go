package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type JwtWrapper struct {
	SecretKey string
	Issuer string
}

type JwtClaim struct {
	Email string
	jwt.StandardClaims
}

func (j *JwtWrapper) GenerateToken(email string) (signedToken string, err error)