/*
 * @Author: gongluck
 * @Date: 2020-06-03 11:42:02
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-13 13:40:05
 */

package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 将结果格式化成json返回
func ReturnResult(c *gin.Context, status int, ret int, data interface{}) {
	c.JSON(status, gin.H{
		"ret":  ret,
		"data": data,
	})
}

// 用户注册
func ApiRegist(c *gin.Context) {
	user, statuscode, err := regist(c)

	if err != nil {
		ReturnResult(c, statuscode, -1, err.Error())
		return
	}

	fmt.Println("user", user.Name, "regist succeed.", user)
	ReturnResult(c, statuscode, 0, "regist succeed.")
	return
}

// 用户登录
func ApiLogin(c *gin.Context) {
	user, statuscode, err := login(c)

	if err != nil {
		ReturnResult(c, statuscode, -1, err.Error())
		return
	}

	fmt.Println("user", user.Name, "login succeed.", user)
	ReturnResult(c, statuscode, 0, "login succeed.")
	return
}

// 用户注销
func ApiLogout(c *gin.Context) {
	_, statuscode, err := logout(c)

	if err != nil {
		ReturnResult(c, statuscode, -1, err.Error())
		return
	}

	fmt.Println("user", "logout succeed.")
	ReturnResult(c, statuscode, 0, "logout succeed.")
	return
}

// 获取所有视频信息
func ApiGetVideos(c *gin.Context) {
	videos, statuscode, err := getvideos(c)

	if err != nil {
		ReturnResult(c, statuscode, -1, err.Error())
		return
	}

	var data []byte
	data, err = json.Marshal(videos)
	if err != nil {
		fmt.Println("JSON marshaling failed:", err)
	} else {
		fmt.Printf("%s\n", data)
	}

	ReturnResult(c, http.StatusOK, 0, string(data))
	return
}

// 上传视频
func ApiPostVideo(c *gin.Context) {
	video, statuscode, err := postvideo(c)

	if err != nil {
		ReturnResult(c, statuscode, -1, err.Error())
		return
	}
	fmt.Println("post video succeed.", video)

	var data []byte
	data, err = json.Marshal(video)
	if err != nil {
		fmt.Println("JSON marshaling failed:", err)
	} else {
		fmt.Printf("%s\n", data)
	}

	ReturnResult(c, http.StatusOK, 0, string(data))
	return
}

// 删除视频
func ApiDelVideo(c *gin.Context) {
	video, statuscode, err := delvideo(c)

	if err != nil {
		ReturnResult(c, statuscode, -1, err.Error())
		return
	}
	fmt.Println("delete video succeed.", video)

	var data []byte
	data, err = json.Marshal(video)
	if err != nil {
		fmt.Println("JSON marshaling failed:", err)
	} else {
		fmt.Printf("%s\n", data)
	}

	ReturnResult(c, http.StatusOK, 0, string(data))
	return
}
