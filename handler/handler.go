/*
 * @Author: gongluck
 * @Date: 2020-06-13 12:32:52
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-17 15:48:41
 */

package handler

import (
	"errors"
	"fmt"
	"govideo_server/conf"
	"govideo_server/dao"
	"govideo_server/model"
	"govideo_server/util"
	"net/http"
	"os"
	"path/filepath"
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
		Level:    int64(dao.GetUsersCount()),
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
func logout(c *gin.Context) (int64, int, error) {
	userid, err := util.DelSession(c)
	return userid, http.StatusOK, err
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
	if file.Size > conf.Config.Video.MaxFileSize {
		return nil, http.StatusInternalServerError, errors.New("can not post file large than " + strconv.FormatInt(conf.Config.Video.MaxFileSize, 10))
	}

	newfilename := util.NewUUID() + ".mp4"

	// create videos directory if it doesn't exist
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	videoDir := filepath.Join(cwd, conf.Config.Video.FilePrefix)
	if _, err := os.Stat(videoDir); os.IsNotExist(err) {
		os.Mkdir(videoDir, 0755)
	}

	err = c.SaveUploadedFile(file, conf.Config.Video.FilePrefix+newfilename)
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

// 删除视频
func delvideo(c *gin.Context) (*model.Video, int, error) {
	postuserid := util.GetUserID(c)
	if postuserid == 0 {
		return nil, http.StatusBadRequest, errors.New("please login first.")
	}
	postuser := dao.GetUserByID(postuserid)
	if postuser == nil {
		return nil, http.StatusBadRequest, errors.New("wrong user.")
	}

	videoid := c.PostForm("videoid")
	if videoid == "" {
		return nil, http.StatusBadRequest, errors.New("your videoid is wrong.")
	}
	vid, err := strconv.ParseInt(videoid, 10, 64)
	if err != nil {
		return nil, http.StatusBadRequest, errors.New("your videoid is wrong." + err.Error())
	}
	video := dao.GetVideoByID(vid)
	if video.ID == 0 {
		return nil, http.StatusBadRequest, errors.New("can not find video by id " + strconv.FormatInt(vid, 10))
	}

	videouser := dao.GetUserByID(video.Userid)
	if videouser == nil || videouser.ID == postuser.ID || videouser.Level > postuser.Level {
		fmt.Println("videouser:", videouser)
		fmt.Println("postuser:", postuser)
		dao.DelVideo(video)
		return video, http.StatusOK, nil
	} else {
		return nil, http.StatusBadRequest, errors.New("you can not delete video.")
	}
}
