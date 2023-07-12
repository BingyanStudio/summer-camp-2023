package controller

import (
	"log"
	"net/http"
	"system/app/shared/database"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	userID := c.Query("id")
	if database.Checkconnection() {
		log.Println("数据库连接成功,即将进行删除操作")
		log.Println("删除用户ID: ", userID)
	}

	// 删除用户
	stmt, err := database.Db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to prepare statement"})
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
