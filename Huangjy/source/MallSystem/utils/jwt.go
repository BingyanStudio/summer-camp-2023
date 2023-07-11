package utils

import (
	"MallSystem/model/response"
	"net/http"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func GenerateJWTToken(id string) string {
	// 这里jwtKey不会为空
	jwtKey := viper.GetString("JWT.SecretKey")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": id,
	})
	// 这里不会有error被返回
	tokenString, _ := token.SignedString([]byte(jwtKey))
	return tokenString
}

func MiddlewareJWTAuthorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if len(auth) <= len("Bearer ") {
			c.JSON(http.StatusUnauthorized, response.UnautherizedError)
			c.Abort()
			return
		}
		tokenString := auth[len("Bearer "):]
		token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return viper.GetString("JWT.SecretKey"), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, response.UnautherizedError)
			c.Abort()
			return
		}
		claims, ok := token.Claims.(*jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, response.UnautherizedError)
			c.Abort()
			return
		}
		c.Set("userid", (*claims)["userid"].(string))
	}
}
