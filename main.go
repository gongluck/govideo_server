/*
 * @Author: gongluck
 * @Date: 2020-06-03 11:34:06
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-04 11:20:19
 */

package main

import (
	"govideo_server/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api")
	{
		api.GET("/getvideos", handler.GetVideos)

		api.POST("/postvideo", handler.PostVideo)
	}

	web := r.Group("/web")
	{
		web.GET("/getvideos", handler.WebGetVideos)
	}

	r.StaticFS("/videos", http.Dir("./videos"))
	r.LoadHTMLGlob("templates/*")

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
