/*
 * @Author: gongluck
 * @Date: 2020-06-03 11:34:06
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-10 09:56:25
 */

package main

import (
	"govideo_server/defs"
	"govideo_server/handler"
	"govideo_server/util"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	// Gin引擎
	r := gin.Default()

	// Redis连接
	store, err := redis.NewStore(defs.RedisConnSize, defs.RedisNetWork, defs.RedisAddress, defs.RedisPassword, []byte(defs.RedisKey))
	if err != nil {
		panic("failed to connect redis")
	}
	// 随机生成前缀
	redis.SetKeyPrefix(store, util.NewUUID())
	// 使用session和redis
	r.Use(sessions.Sessions(defs.SessionName, store))

	// ping-pong接口，测试网络
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// API路由
	api := r.Group("/api")
	{
		api.POST("/regist", handler.Regist)
		api.POST("/login", handler.Login)
		api.POST("/logout", handler.Logout)

		api.POST("/getvideos", handler.GetVideos)
		api.POST("/postvideo", handler.PostVideo)
	}

	// WEB路由
	web := r.Group("/web")
	{
		web.GET("/getvideos", handler.WebGetVideos)
		web.GET("/upload", handler.WebUploadVideo)
	}

	// 静态文件服务，获取视频文件
	r.StaticFS("/videos", http.Dir("./videos"))

	// 设置模板路径
	r.LoadHTMLGlob("templates/*")

	// 启动HTTP服务
	r.Run(defs.HttpAddr)
}
