/*
 * @Author: gongluck
 * @Date: 2020-06-03 11:34:06
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-03 15:23:38
 */

package main

import (
	"govideo_server/handler"

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

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
