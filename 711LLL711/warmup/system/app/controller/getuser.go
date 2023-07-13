package controller

import (
	"database/sql"
	"log"
	"net/http"
	"system/app/shared/database"
	myuser "system/app/shared/user"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userID := c.Query("id")
	log.Println("查询用户ID: ", userID)

	var user myuser.User
	err := database.Db.QueryRow("SELECT * FROM users WHERE id = ?", userID).Scan(
		&user.Id,
		&user.Password,
		&user.Username,
		&user.Phone,
		&user.Email,
		&user.Role,
	)
	if err != nil {
		log.Println("查询用户失败")
		if err == sql.ErrNoRows {
			log.Println("用户不存在")
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	} else {
		log.Println("查询用户成功")
	}
	c.JSON(http.StatusOK, user)
}

// 获取所有用户
func GetAllUsers(c *gin.Context) {
	rows, err := database.Db.Query("SELECT * FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}
	defer rows.Close()

	var users []myuser.User
	for rows.Next() {
		var user myuser.User
		scanErr := rows.Scan(
			&user.Id,
			&user.Password,
			&user.Username,
			&user.Phone,
			&user.Email,
			&user.Role,
		)
		if scanErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan user"})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}
