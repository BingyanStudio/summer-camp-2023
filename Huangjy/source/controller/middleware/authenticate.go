package middleware

import (
	db "myserver/database"
	"myserver/model"
	"net/http"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

var jwtKey []byte = []byte(db.QueryOtherInfo("jwtKey").(string))

func Authorizate(role_needed uint8) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"succeed": false,
				"error":   "Missing authorization",
				"data":    "",
			})
			c.Abort()
		}
		tokenString := authHeader[len("Bearer "):]
		token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"succeed": false,
				"error":   "Sorry, we come across some problem",
				"data":    "",
			})
			c.Abort()
		}
		claims, ok := token.Claims.(*jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"succeed": false,
				"error":   "Invalid token",
				"data":    "",
			})
			c.Abort()
		}
		role := uint8((*claims)["role"].(float64))
		id := uint32((*claims)["userid"].(float64))
		if role < role_needed {
			c.JSON(http.StatusUnauthorized, gin.H{
				"succeed": false,
				"error":   "Unauthorized",
				"data":    "",
			})
			c.Abort()
		}
		c.Set("curUser", &model.UserInfo{
			ID:   id,
			Role: role,
		})
	}
}
