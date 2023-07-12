package controller

import (
	"log"
	"net/http"
	"system/app/shared/database"
	myuser "system/app/shared/user"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Updateuser(user myuser.User, c *gin.Context) {
	//生成哈希密码
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	// 更新用户信息到数据库
	if database.Checkconnection() {
		log.Println("数据库连接成功,即将进行更新操作")
	}
	stmt, err := database.Db.Prepare("UPDATE users SET name = ?, phone = ?, email = ? ,password = ?,phone  = ? ,role = ? WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to prepare statement"})
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Phone, user.Email, hashedPassword, user.Phone, user.Role, user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
