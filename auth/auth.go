package auth

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

const (
	UserId    = "UserId"
	UserEmail = "Email"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
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

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatal(err.Error())
	}

	verifyBytes, err := ioutil.ReadFile(os.Getenv("AUTH_JWT_PUBLIC_KEY_PATH"))
	if err != nil {
		log.Fatalf("Failed to read %s", os.Getenv("AUTH_JWT_PUBLIC_KEY_PATH"))
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
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

	return token.SignedString(signKey)
}

type AuthHeader struct {
	Header string `header:"Authorization"`
}

func AuthenticateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := AuthHeader{}
		if err := c.ShouldBindHeader(&header); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()

			return
		}

		token, err := request.ParseFromRequestWithClaims(c.Request, request.OAuth2Extractor, &JwtClaim{}, func(t *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid Token",
			})
			c.Abort()

			return
		}

		c.Set(UserId, token.Claims.(*JwtClaim).ID)
		c.Set(UserEmail, token.Claims.(*JwtClaim).Email)

		c.Next()
	}
}
