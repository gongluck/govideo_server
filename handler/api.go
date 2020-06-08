/*
 * @Author: gongluck
 * @Date: 2020-06-03 11:42:02
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-08 15:48:39
 */

package handler

import (
	"encoding/json"
	"fmt"
	"govideo_server/dao"
	"govideo_server/model"
	"govideo_server/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ReturnResult(c *gin.Context, status int, ret int, data interface{}) {
	c.JSON(status, gin.H{
		"ret":  ret,
		"data": data,
	})
}

func Regist(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	if name == "" || password == "" {
		ReturnResult(c, http.StatusBadRequest, -1, "name or password wrong.")
		return
	}

	user := dao.GetUserByName(name)
	if user.ID != 0 {
		ReturnResult(c, http.StatusBadRequest, -1, "you can not use this username "+name)
		return
	}

	if !dao.AddUser(&model.User{
		Name:     name,
		Password: password,
		Level:    10,
	}) {
		ReturnResult(c, http.StatusInternalServerError, -1, "regist fail.")
	}

	ReturnResult(c, http.StatusOK, 0, "login succeed.")
}

func Login(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	if name == "" || password == "" {
		ReturnResult(c, http.StatusBadRequest, -1, "name or password wrong.")
		return
	}

	user := dao.GetUserByName(name)
	if user.ID == 0 {
		ReturnResult(c, http.StatusBadRequest, -1, "can not find user "+name)
		return
	} else if user.Password != password {
		ReturnResult(c, http.StatusBadRequest, -1, "password wrong.")
		return
	}

	err := util.SetSession(c, user.ID)
	if err != nil {
		ReturnResult(c, http.StatusInternalServerError, -1, "fail to create session.")
		return
	}

	ReturnResult(c, http.StatusOK, 0, "login succeed.")
}

func Logout(c *gin.Context) {
	util.DelSession(c)
	ReturnResult(c, http.StatusOK, 0, "logout succeed.")
}

func GetVideos(c *gin.Context) {
	user := util.GetSessionUser(c)
	if user == 0 {
		//ReturnResult(c, http.StatusNotAcceptable, -1, "Please login first.")
		//return
	}

	videos := dao.GetVideos()
	data, err := json.Marshal(videos)
	if err != nil {
		fmt.Println("JSON marshaling failed: ", err)
	} else {
		fmt.Printf("%s\n", data)
	}

	ReturnResult(c, http.StatusOK, 0, string(data))
}

func PostVideo(c *gin.Context) {
	user := util.GetSessionUser(c)
	if user == 0 {
		//匿名用户
	}

	title := c.PostForm("title")
	description := c.PostForm("description")
	if title == "" {
		ReturnResult(c, http.StatusBadRequest, -1, "your title is wrong.")
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		ReturnResult(c, http.StatusInternalServerError, -1, "can not read post file param.")
		return
	}
	if file.Size > 50*1024*1024 {
		ReturnResult(c, http.StatusInternalServerError, -1, "can not post file large than "+strconv.Itoa(50*1024*1024))
		return
	}

	newfilename := "videos/" + util.NewUUID() + ".mp4"
	err = c.SaveUploadedFile(file, newfilename)
	if err != nil {
		ReturnResult(c, http.StatusInternalServerError, -1, "save file failed.")
		return
	} else {
		video := &model.Video{
			Title:       title,
			Description: description,
			Filepath:    newfilename,
			Userid:      user,
		}
		if !dao.AddVideo(video) {
			ReturnResult(c, http.StatusInternalServerError, -1, "insert database failed.")
			return
		} else {
			ReturnResult(c, http.StatusOK, 0, "post succeed.")
			return
		}
	}
}
