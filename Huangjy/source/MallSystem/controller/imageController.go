package controller

import (
	"MallSystem/model/response"
	"bytes"
	"image/png"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func GetImageHandler(c *gin.Context) {
	filename := c.Param("filename")
	if filename == "" {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}

	storagePath := viper.GetString("Image.Source")
	targetPath := filepath.Join(storagePath, filename)
	file, err := os.Open(targetPath)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse("图片未找到"))
		return
	}
	defer file.Close()

	c.Header("Content-Type", "image/"+filepath.Ext(filename)[1:])
	c.File(targetPath)
}

func GetSmallImageHandler(c *gin.Context) {
	filename := c.Param("filename")
	if filename == "" {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}

	storagePath := viper.GetString("Image.Source")
	targetPath := filepath.Join(storagePath, filename)

	img, err := imaging.Open(targetPath)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse("图片未找到"))
		return
	}

	width := 200
	buf := new(bytes.Buffer)
	dstImg := imaging.Resize(img, width, 0, imaging.Lanczos)

	err = png.Encode(buf, dstImg)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/image/"+filename)
		return
	}
	c.Header("Content-Type", "image/"+filepath.Ext(filename)[1:])
	c.Header("Content-Length", strconv.FormatInt(int64(buf.Len()), 10))

	c.Writer.Write(buf.Bytes())
}
