package main

import (
	//"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 自定义中间件（拦截器）for logging, authentication, error handling...
func myMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("name", "张三")
		//这里可以加条件判断是否拦截，比如登录验证是否通过...
		if context.Query("user") == "admin" {
			context.Next() //放行
		} else {
			context.Abort() //拦截
		}
	}
}

func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "小王子") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序
		// c.Abort()
		// 计算耗时
		cost := time.Since(start)
		log.Println(cost)
	}
}

// 给前端响应数据或页面
func main() {
	//创建一个服务
	ginserver := gin.Default()

	ginserver.Use(myMiddleware(), StatCost()) //使用中间件

	//加载静态页面
	ginserver.LoadHTMLGlob("templates/*")

	//加载资源文件
	ginserver.Static("/static", "./static")

	//给前端响应数据或页面
	ginserver.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"title": "首页",
			"msg":   "后台传来的数据",
		})
	})

	//获取前端上传的数据
	ginserver.POST("/upload", func(context *gin.Context) {
		file, err := context.FormFile("file")
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"code": 0,
				"msg":  "上传失败",
			})
			return
		}
		//上传文件到本地目录
		log.Println(file.Filename)
		dst := fmt.Sprintf("C:/tmp/%s", file.Filename)
		context.SaveUploadedFile(file, dst)
		context.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "上传成功",
		})
	})

	//404 router
	ginserver.NoRoute(func(context *gin.Context) {
		context.HTML(http.StatusNotFound, "404.html", gin.H{
			"title": "404",
			"msg":   "页面找不到了",
		})
	})

	//http重定向
	ginserver.GET("/redirect", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	//路由重定向
	ginserver.GET("/redirect2", func(context *gin.Context) {
		context.Request.URL.Path = "/index"
		ginserver.HandleContext(context)
	})

	//获取请求中的参数
	//usl?user=xxx&pwd=xxx
	ginserver.GET("/user", myMiddleware(), func(context *gin.Context) {
		//取出中间件的值
		user := context.MustGet("name").(string)
		log.Println("v ...any", user)
		username := context.Query("user")
		pwd := context.Query("pwd")
		fmt.Println(username, pwd)
		context.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "请求成功",
			"data": gin.H{
				"username": username,
				"pwd":      pwd,
			},
		})
	})

	///user/xxx/xxx 获取path参数
	ginserver.GET("/user/:name/:pwd", func(context *gin.Context) {
		name := context.Param("name")
		pwd := context.Param("pwd")
		fmt.Println(name, pwd)
		context.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "请求成功",
			"data": gin.H{
				"name": name,
				"pwd":  pwd,
			},
		})
	})

	//获取form参数
	ginserver.POST("/form", func(context *gin.Context) {
		username := context.PostForm("username")
		pwd := context.PostForm("pwd")
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"user":    username,
			"pwd":     pwd,
		})
	})

	//使用路由组,用大括号包裹
	//group := ginserver.Group("/admin")

	//服务器端口
	ginserver.Run(":8080")
}
