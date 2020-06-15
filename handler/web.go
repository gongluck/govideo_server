/*
 * @Author: gongluck
 * @Date: 2020-06-04 09:48:42
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-13 14:02:27
 */
package handler

import (
	"fmt"
	"govideo_server/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// web主页
func WebIndex(c *gin.Context) {
	videos, statuscode, _ := getvideos(c)
	c.HTML(statuscode, "index.html", videos)
	return
}

// 获取所有视频信息
func WebGetVideos(c *gin.Context) {
	videos, statuscode, _ := getvideos(c)
	c.HTML(statuscode, "index.html", videos)
	return
}

// 上传视频
func WebPostVideo(c *gin.Context) {
	c.HTML(http.StatusOK, "postvideo.html", nil)
	return
}

// 上传结果
func WebPostVideoResult(c *gin.Context) {
	video, statuscode, err := postvideo(c)
	if err != nil {
		fmt.Println("upload failed,", err.Error())
		c.HTML(statuscode, "postvideoresult.html", &model.Video{
			ID: 0,
		})
	} else {
		c.HTML(statuscode, "postvideoresult.html", video)
		return
	}
}
