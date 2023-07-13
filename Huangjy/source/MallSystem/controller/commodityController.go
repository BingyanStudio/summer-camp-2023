package controller

import (
	"MallSystem/database"
	"MallSystem/model"
	"MallSystem/model/response"
	"MallSystem/utils"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetCommoditiesHandler(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		page = 0
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil || limit > 20 || limit < 0 {
		limit = 20
	}
	category, err := strconv.Atoi(c.DefaultQuery("category", "0"))
	if err != nil || category < int(model.All) || category > int(model.Others) {
		category = int(model.All)
	}
	keyword := c.Query("keyword")
	filter := bson.M{"title": primitive.Regex{Pattern: keyword, Options: "i"}, "status": model.Selling}
	opt := options.Find().SetSkip(int64(page * limit)).SetLimit(int64(limit))
	results, err := database.QueryCommodities(&filter, opt)
	infos := make([]model.CommoditySearchResult, 0)
	if err != nil {
		if err == context.DeadlineExceeded {
			c.JSON(http.StatusInternalServerError, response.TimeoutError)
		} else {
			c.JSON(http.StatusBadRequest, response.MakeFailedResponse(err.Error()))
		}
		return
	}
	for _, r := range results {
		c := model.CommoditySearchResult{}
		c.Category = r.Category
		c.ID = r.ID
		c.Picture = r.Picture
		c.Price = r.Price
		c.Title = r.Title
		infos = append(infos, c)
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(infos))
}

func GetOneCommodityHandler(c *gin.Context) {
	commodityID, err := primitive.ObjectIDFromHex(c.Param("commodityID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse("没有找到商品"))
	}
	ci, err := database.QueryOneCommodity(&bson.M{"_id": commodityID})
	if err != nil {
		if err == context.DeadlineExceeded {
			c.JSON(http.StatusInternalServerError, response.TimeoutError)
		} else {
			c.JSON(http.StatusBadRequest, response.MakeFailedResponse("没有找到商品"))
		}
		return
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(*ci))
}

func GetHotKeywordHandler(c *gin.Context) {

}

func PubCommodityHandler(c *gin.Context) {
	var (
		ci model.CommodityInfo
	)
	if err := c.ShouldBind(&ci); err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	if err := utils.ValidateCommodityInfo(&ci); err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse(err.Error()))
		return
	}
	ci.UserID, _ = primitive.ObjectIDFromHex(c.GetString("userid")[10:34])
	ci.LaunchTime = time.Now()
	ci.Status = model.Selling
	if err := database.InsertOneCommodity(&ci); err != nil {
		if err == context.DeadlineExceeded {
			c.JSON(http.StatusInternalServerError, response.TimeoutError)
		} else {
			c.JSON(http.StatusBadRequest, response.MakeFailedResponse(err.Error()))
		}
		return
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(""))
}
