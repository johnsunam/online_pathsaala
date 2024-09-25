package middleware

import (
	"fmt"
	"net/http"
	"online-pathsaala/model"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		tokenArr := strings.SplitAfter(tokenString, " ")
		if len(tokenArr) < 2 {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "Invalid token format")
		}
		claims := model.SignedDetails{}
		token, err := jwt.ParseWithClaims(tokenArr[1], &claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			SECRET_KEY := os.Getenv("SECRET_KEY")

			return []byte(SECRET_KEY), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "token expired")
			return
		}
		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "token expired")
			return
		}
		c.Set("userId", claims.Id)
		c.Next()
	}
}
