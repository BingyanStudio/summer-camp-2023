package controller

import (
	"MallSystem/database"
	"MallSystem/model"
	"MallSystem/model/response"
	"MallSystem/utils"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func UserRegisterHandler(c *gin.Context) {
	var (
		u model.UserInfo
	)
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	if err := utils.ValidateRegisterInfo(&u); err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse(err.Error()))
		return
	}
	if err := utils.EncryptUserPassword(&u.Password); err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	if err := database.InsertOneUser(&u); err != nil {
		if err == context.DeadlineExceeded {
			c.JSON(http.StatusInternalServerError, response.TimeoutError)
		} else {
			c.JSON(http.StatusBadRequest, response.MakeFailedResponse("用户名已存在"))
		}
		return
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(""))
}

func UserLoginHandler(c *gin.Context) {
	var (
		l model.Login
		u *model.UserInfo
	)
	if err := c.ShouldBind(&l); err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	u, err := database.QueryOneUser(&bson.M{
		"username": l.Username,
	})
	if err != nil {
		if err == context.DeadlineExceeded {
			c.JSON(http.StatusInternalServerError, response.TimeoutError)
		} else {
			c.JSON(http.StatusBadRequest, response.MakeFailedResponse("用户名错误"))
		}
		return
	}
	if err := utils.ComparePassword(&l, u); err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse("密码错误"))
		return
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(utils.GenerateJWTToken(u.ID.String())))
}
