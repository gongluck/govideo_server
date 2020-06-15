/*
 * @Author: gongluck
 * @Date: 2020-06-13 12:32:52
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-13 13:32:04
 */

package handler

import (
	"errors"
	"govideo_server/dao"
	"govideo_server/defs"
	"govideo_server/model"
	"govideo_server/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 用户注册
func regist(c *gin.Context) (*model.User, int, error) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	if name == "" || password == "" {
		return nil, http.StatusBadRequest, errors.New("name or password wrong.")
	}

	user := dao.GetUserByName(name)
	if user.ID != 0 {
		return nil, http.StatusBadRequest, errors.New("you can not use this username " + name)
	}

	user = &model.User{
		Name:     name,
		Password: password,
		Level:    10,
	}
	if !dao.AddUser(user) {
		return nil, http.StatusInternalServerError, errors.New("regist fail.")
	}

	return user, http.StatusOK, nil
}

// 用户登录
func login(c *gin.Context) (*model.User, int, error) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	if name == "" || password == "" {
		return nil, http.StatusBadRequest, errors.New("name or password wrong.")
	}

	user := dao.GetUserByName(name)
	if user.ID == 0 {
		return nil, http.StatusBadRequest, errors.New("can not find user " + name)
	} else if user.Password != password {
		return nil, http.StatusBadRequest, errors.New("password wrong.")
	}

	err := util.SetSession(c, user.ID)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("fail to create session.")
	}

	return user, http.StatusOK, nil
}

// 用户注销
func logout(c *gin.Context) (*model.User, int, error) {
	util.DelSession(c)
	return nil, http.StatusOK, nil
}

// 获取所有视频信息
func getvideos(c *gin.Context) ([]*model.Video, int, error) {
	videos := dao.GetVideos()
	return videos, http.StatusOK, nil
}

// 上传视频
func postvideo(c *gin.Context) (*model.Video, int, error) {
	userid := util.GetUserID(c)
	title := c.PostForm("title")
	description := c.PostForm("description")

	if title == "" {
		return nil, http.StatusBadRequest, errors.New("your title is wrong.")
	}

	file, err := c.FormFile("file")
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("can not read post file param.")
	}
	if file.Size > defs.MaxFileSize {
		return nil, http.StatusInternalServerError, errors.New("can not post file large than " + strconv.FormatInt(defs.MaxFileSize, 10))
	}

	newfilename := util.NewUUID() + ".mp4"
	err = c.SaveUploadedFile(file, defs.FilePrefix+newfilename)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("save file failed," + newfilename)
	} else {
		video := &model.Video{
			Title:       title,
			Description: description,
			Filepath:    "videos/" + newfilename,
			Userid:      userid,
		}
		if !dao.AddVideo(video) {
			return nil, http.StatusInternalServerError, errors.New("insert database failed.")
		} else {
			return video, http.StatusOK, nil
		}
	}
}
