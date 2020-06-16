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

// 用户注册
func WebRegistPage(c *gin.Context) {
	c.HTML(http.StatusOK, "regist.html", nil)
	return
}
func WebRegist(c *gin.Context) {
	user, statuscode, err := regist(c)

	if err != nil {
		fmt.Println("regist failed,", err.Error())
		c.HTML(statuscode, "videoresult.html", &model.User{
			ID: 0,
		})
	} else {
		c.HTML(statuscode, "videoresult.html", user)
		return
	}
}

// 用户登录
func WebLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
	return
}
func WebLogin(c *gin.Context) {
	user, statuscode, err := login(c)

	if err != nil {
		fmt.Println("login failed,", err.Error())
		c.HTML(statuscode, "videoresult.html", &model.User{
			ID: 0,
		})
	} else {
		c.HTML(statuscode, "videoresult.html", user)
		return
	}
	return
}

// 用户注销
func WebLogout(c *gin.Context) {
	user, statuscode, err := logout(c)

	if err != nil {
		fmt.Println("logout failed,", err.Error())
		c.HTML(statuscode, "videoresult.html", &model.User{
			ID: 0,
		})
	} else {
		c.HTML(statuscode, "videoresult.html", user)
		return
	}
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
		c.HTML(statuscode, "videoresult.html", &model.Video{
			ID: 0,
		})
	} else {
		c.HTML(statuscode, "videoresult.html", video)
		return
	}
}

// 删除视频
func WebDelVideo(c *gin.Context) {
	video, statuscode, err := delvideo(c)
	if err != nil {
		fmt.Println("delete failed,", err.Error())
		c.HTML(statuscode, "videoresult.html", &model.Video{
			ID: 0,
		})
	} else {
		c.HTML(statuscode, "videoresult.html", video)
		return
	}
}
