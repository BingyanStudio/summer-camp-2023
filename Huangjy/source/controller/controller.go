package controller

import (
	"encoding/json"
	"log"
	"myserver/controller/utils"
	db "myserver/database"
	"myserver/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(c *gin.Context) {
	reJSON := func(code int, s bool, err, d string) {
		c.JSON(code, gin.H{
			"succeed": s,
			"error":   err,
			"data":    d,
		})
	}
	log.Println("user register...")
	u := model.UserInfo{Role: 1}
	if err := c.ShouldBind(&u); err != nil {
		log.Println(err)
		reJSON(http.StatusBadRequest, false, "Invalid info", "")
		return
	}

	if err := utils.CheckRegisterValid(&u); err != nil {
		log.Println(err)
		reJSON(http.StatusBadRequest, false, err.Error(), "")
		return
	}

	if res, err := db.QueryUser(&bson.D{{"name", u.Name}}, options.Find().SetLimit(1)); err != nil {
		log.Println(err)
		reJSON(http.StatusInternalServerError, false, "Sorry, we cross over an error", "")
		return
	} else if len(res) != 0 {
		reJSON(http.StatusBadRequest, false, "Name has been already used", "")
		return
	}

	hashed_pwd, err := bcrypt.GenerateFromPassword([]byte(u.Pwd), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		reJSON(http.StatusInternalServerError, false, "Sorry, we cross over an error", "")
		return
	}
	u.Pwd = string(hashed_pwd)
	if err := db.InsertUser(&u); err != nil {
		log.Println(err)
		reJSON(http.StatusInternalServerError, false, "Sorry, we cross over an error", "")
		return
	}
	reJSON(http.StatusOK, true, "", "Register successful")
}

func LoginHandler(c *gin.Context) {
	reJSON := func(code int, s bool, err, d string) {
		c.JSON(code, gin.H{
			"succeed": s,
			"error":   err,
			"data":    d,
		})
	}
	log.Println("user login...")
	var login_form model.UserLogin

	if err := c.ShouldBind(&login_form); err != nil {
		log.Println(err)
		reJSON(http.StatusBadRequest, false, "Invalid info", "")
		return
	}
	if login_form.Username == "" {
		reJSON(http.StatusBadRequest, false, "Name could not be empty", "")
		return
	}

	u, err := db.QueryUser(&bson.D{{"name", login_form.Username}}, options.Find().SetLimit(1))
	if err != nil {
		log.Println(err)
		reJSON(http.StatusInternalServerError, false, "Sorry, we cross over an error", "")
		return
	} else if len(u) == 0 {
		reJSON(http.StatusBadRequest, false, "Name does not exist", "")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u[0].Pwd), []byte(login_form.Password)); err != nil {
		reJSON(http.StatusBadRequest, false, "Password is wrong", "")
		return
	}
	token := utils.GenerateJWTToken(u[0].ID, u[0].Role)
	if token == "" {
		log.Println("token is empty")
		reJSON(http.StatusInternalServerError, false, "Sorry, we cross over an error", "")
		return
	}
	reJSON(http.StatusOK, true, "", token)
}

func UserUpdateInfoHandler(c *gin.Context) {
	var update model.UserUpdate
	err := c.ShouldBind(&update)
	log.Println(update)
	u, ok := c.Get("curUser")
	log.Println(u, ok, err)
	log.Println("to there 1", err != nil)
	if err != nil || !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"succeed": false,
			"error":   "Invalid info",
			"data":    "",
		})
		return
	}
	log.Println("to there 2")
	if err := db.UpdateUser(&bson.D{{"id", uint32(u.(*model.UserInfo).ID)}}, &update); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"succeed": false,
			"error":   "Sorry, we come across some problems",
			"data":    "",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"succeed": true,
		"error":   "",
		"data":    "Update successful",
	})
}

func AdminDeleteHandler(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil || id > ^uint64(0)>>32 {
		c.JSON(http.StatusBadRequest, gin.H{
			"succeed": false,
			"error":   "Invalid id",
			"data":    "",
		})
		return
	}
	if err := db.DeleteUser(&bson.D{{"id", id}}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"succeed": false,
			"error":   "Invalid id",
			"data":    "",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"succeed": true,
		"error":   "",
		"data":    "Delete successful",
	})
}

func AdminGetOneUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil || id > ^uint64(0)>>32 {
		c.JSON(http.StatusBadRequest, gin.H{
			"succeed": false,
			"error":   "Invalid id",
			"data":    "",
		})
		return
	}
	users, err := db.QueryUser(&bson.D{{"id", id}}, options.Find().SetLimit(1))
	if err != nil || len(users) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"succeed": false,
			"error":   "Invalid id",
			"data":    "",
		})
		return
	}
	users[0].Pwd = ""
	b, err := json.Marshal(*users[0])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"succeed": false,
			"error":   "Sorry, we come across some problems",
			"data":    "",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"succeed": true,
		"error":   "",
		"data":    string(b),
	})
}

func AdminGetAllUsers(c *gin.Context) {
	users, err := db.QueryUser(&bson.D{{}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"succeed": false,
			"error":   "Sorry, we come across some problems",
			"data":    "",
		})
		return
	}
	var usersString string
	for _, u := range users {
		u.Pwd = ""
		b, err := json.Marshal(*users[0])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"succeed": false,
				"error":   "Sorry, we come across some problems",
				"data":    "",
			})
			return
		}
		usersString += string(b) + ", "
	}
	c.JSON(http.StatusOK, gin.H{
		"succeed": true,
		"error":   "",
		"data":    usersString,
	})
}
