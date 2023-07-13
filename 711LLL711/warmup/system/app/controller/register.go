package controller

import (
	"log"
	"net/http"
	"system/app/shared/database"
	myuser "system/app/shared/user"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user myuser.User, c *gin.Context) {
	// 加密密码
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	log.Println("生成哈希密码： ", hashedPassword)
	// 插入用户信息到数据库
	stmt, err := database.Db.Prepare("INSERT INTO users (id,password, name, phone, email, role) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to prepare statement"})
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id, hashedPassword, user.Username, user.Phone, user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
