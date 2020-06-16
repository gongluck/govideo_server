/*
 * @Author: gongluck
 * @Date: 2020-06-03 11:34:06
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-13 14:08:55
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
	gin.SetMode(defs.GinMode)
	r := gin.Default()

	// Redis连接
	store, err := redis.NewStore(defs.RedisConnSize, defs.RedisNetWork, defs.RedisAddress, defs.RedisPassword, []byte(defs.RedisKey))
	if err != nil {
		panic("failed to connect redis " + defs.RedisAddress + " " + err.Error())
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
		api.POST("/regist", handler.ApiRegist)
		api.POST("/login", handler.ApiLogin)
		api.POST("/logout", handler.ApiLogout)

		api.POST("/getvideos", handler.ApiGetVideos)
		api.POST("/postvideo", handler.ApiPostVideo)
		api.POST("/delvideo", handler.ApiDelVideo)
	}

	// WEB路由
	web := r.Group("/web")
	{
		web.GET("/", handler.WebIndex)

		web.GET("/regist", handler.WebRegistPage)
		web.POST("/regist", handler.WebRegist)
		web.GET("/login", handler.WebLoginPage)
		web.POST("/login", handler.WebLogin)
		web.GET("/logout", handler.WebLogout)

		// 视频上传页面
		web.GET("/postvideo", handler.WebPostVideo)
		web.POST("/postvideoresult", handler.WebPostVideoResult)

		// 删除视频
		web.POST("/delvideo", handler.WebDelVideo)
	}

	// 静态文件服务，获取视频文件和网页资源文件
	r.StaticFS("/videos", http.Dir(defs.FilePrefix))
	r.StaticFS("/static", http.Dir(defs.TemplatesPath+"static"))

	// 设置模板路径
	r.LoadHTMLGlob(defs.TemplatesPath + "*.html")

	// 启动HTTP服务
	r.Run(defs.HttpAddr)
}
