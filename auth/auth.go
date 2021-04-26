package auth

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	VerifyKey *rsa.PublicKey
	SignKey   *rsa.PrivateKey
)

type JwtWrapper struct {
	SecretKey string
	Issuer    string
}

type JwtClaim struct {
	ID    uint8
	Email string
	*jwt.StandardClaims
}

func InitAuth() {
	signBytes, err := ioutil.ReadFile(os.Getenv("AUTH_JWT_PRIVATE_KEY_PATH"))
	if err != nil {
		log.Fatalf("Failed to read %s", os.Getenv("AUTH_JWT_PRIVATE_KEY_PATH"))
	}

	SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatal(err.Error())
	}

	verifyBytes, err := ioutil.ReadFile(os.Getenv("AUTH_JWT_PUBLIC_KEY_PATH"))
	if err != nil {
		log.Fatalf("Failed to read %s", os.Getenv("AUTH_JWT_PUBLIC_KEY_PATH"))
	}

	VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (j *JwtWrapper) GenerateToken(userId uint8, email string) (signedToken string, err error) {
	token := jwt.New(jwt.GetSigningMethod("RS256"))

	token.Claims = &JwtClaim{
		userId,
		email,
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 240).Unix(),
		},
	}

	return token.SignedString(SignKey)
}
