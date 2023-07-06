package controller

import (
	"log"
	"myserver/controller/utils"
	db "myserver/database"
	"myserver/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	u := model.User{Role: 1}
	if err := c.ShouldBind(&u); err != nil {
		log.Println(err)
		reJSON(http.StatusBadRequest, false, "Invalid info", "")
		return
	}
	log.Printf("%+v\n", u)

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
	} else if err := db.InsertUser(&u); err != nil {
		log.Println(err)
		reJSON(http.StatusInternalServerError, false, "Sorry, we cross over an error", "")
		return
	}
	reJSON(http.StatusOK, true, "", "")
}

func LoginHandler(c *gin.Context) {
	log.Println("user login...")

}
