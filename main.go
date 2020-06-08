/*
 * @Author: gongluck
 * @Date: 2020-06-03 11:34:06
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-08 15:30:02
 */

package main

import (
	"govideo_server/handler"
	"govideo_server/util"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	if err != nil {
		panic("failed to connect redis")
	}
	redis.SetKeyPrefix(store, util.NewUUID())
	r.Use(sessions.Sessions("session", store))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api")
	{
		api.POST("/regist", handler.Regist)
		api.POST("/login", handler.Login)
		api.POST("/logout", handler.Logout)

		api.POST("/getvideos", handler.GetVideos)

		api.POST("/postvideo", handler.PostVideo)
	}

	web := r.Group("/web")
	{
		web.GET("/getvideos", handler.WebGetVideos)

		web.GET("/upload", handler.WebUploadVideo)
	}

	r.StaticFS("/videos", http.Dir("./videos"))
	r.LoadHTMLGlob("templates/*")

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
