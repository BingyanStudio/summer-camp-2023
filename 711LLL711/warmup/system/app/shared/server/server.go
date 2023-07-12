package server

//感觉这个server包没有必要，直接放在main.go里面就可以了
import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var Ginserver *gin.Engine

func init() {
	Ginserver = gin.Default()
}
