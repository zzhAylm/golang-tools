package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func filterHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("userSession", "userId-1")

	}
}

func main() {

	ginServer := gin.Default()

	// 加载静态页面
	ginServer.LoadHTMLGlob("templates/*")

	// 加载资源文件
	ginServer.Static("/static", "./static")

	ginServer.GET("/get", func(context *gin.Context) {

		context.JSON(200, gin.H{"msg": "hello world!"})

	})

	ginServer.GET("/user", func(context *gin.Context) {
		userId := context.Query("userId")
		userName := context.Query("userName")
		context.JSON(200, gin.H{"userId": userId, "username": userName})
	})

	ginServer.GET("/index", func(context *gin.Context) {
		context.HTML(200, "index.html", gin.H{"msg": "zzh"})
	})

	ginServer.GET("/user/:userId/:userName", func(context *gin.Context) {
		userId := context.Param("userId")
		userName := context.Param("userName")
		context.JSON(200, gin.H{"userId": userId, "username": userName})
	})

	ginServer.POST("/post", func(context *gin.Context) {
		data, err := context.GetRawData()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"msg": "error"})
		}

		var obj = make(map[string]interface{})

		err = json.Unmarshal(data, &obj)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"msg": "error"})
		}

		context.JSON(http.StatusOK, obj)
	})

	ginServer.POST("/form", func(context *gin.Context) {
		userName := context.PostForm("userName")
		userId := context.PostForm("userId")
		context.JSON(http.StatusOK, gin.H{"userName": userName, "userId": userId})
	})

	// 路由，重定向
	ginServer.GET("/redirect", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	ginServer.NoRoute(func(context *gin.Context) {
		context.HTML(http.StatusNotFound, "404.html", nil)
	})

	userGroup := ginServer.Group("user")
	{
		userGroup.GET("/")
		userGroup.POST("/")
		userGroup.PUT("/")

	}

	ginServer.GET("/login", filterHandler(), func(context *gin.Context) {

		userSession := context.MustGet("userSession").(string)

		context.JSON(http.StatusOK, gin.H{"userSession": userSession})
	})

	ginServer.Run(":8080")
}
