package middleware

import (
	"car_scraper/auth"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	UserId    = "UserId"
	UserEmail = "Email"
)

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

		token, err := request.ParseFromRequest(
			c.Request,
			request.OAuth2Extractor,
			func(t *jwt.Token) (interface{}, error) {
				return auth.VerifyKey, nil
			},
			request.WithClaims(&auth.JwtClaim{}),
		)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid Token",
			})
			c.Abort()

			return
		}

		c.Set(UserId, token.Claims.(*auth.JwtClaim).ID)
		c.Set(UserEmail, token.Claims.(*auth.JwtClaim).Email)

		c.Next()
	}
}
