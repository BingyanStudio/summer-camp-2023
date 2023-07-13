package controller

import (
	"net/http"
	"system/app/shared/database"
	mysession "system/app/shared/session"
	myuser "system/app/shared/user"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 验证密码
func verifyPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func LoginUser(user myuser.User, c *gin.Context) {
	// 查询用户信息
	var retrievedUser myuser.User

	err := database.Db.QueryRow("SELECT password,role FROM users WHERE id = ?", user.Id).Scan(
		&retrievedUser.Password,
		&retrievedUser.Role,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
	}

	// 验证密码
	result := verifyPassword(user.Password, retrievedUser.Password)
	if result {
		// 密码正确·--设置session
		mysession.SetSession(user.Id, retrievedUser.Role, c.Request, c.Writer)

		c.JSON(http.StatusOK, gin.H{"message": "login successful"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "login failed"})
	}
}
