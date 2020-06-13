/*
 * @Author: gongluck
 * @Date: 2020-06-04 09:48:42
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-10 15:54:34
 */
package handler

import (
	"govideo_server/dao"
	"govideo_server/model"
	"govideo_server/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func WebIndex(c *gin.Context) {
	videos := dao.GetVideos()
	c.HTML(http.StatusOK, "index.html", videos)
	return
}

func WebGetVideos(c *gin.Context) {
	videos := dao.GetVideos()
	c.HTML(http.StatusOK, "index.html", videos)
	return
}

func WebUpload(c *gin.Context) {
	videos := dao.GetVideos()
	c.HTML(http.StatusOK, "upload.html", videos)
	return
}

func WebUploadVideo(c *gin.Context) {
	user := util.GetUserID(c)
	if user == 0 {
		//匿名用户
	}

	res := model.UploadResult{}

	title := c.PostForm("title")
	description := c.PostForm("description")
	if title == "" {
		res.Result = "上传失败！"
		res.Desc = "your title is wrong."
		c.HTML(http.StatusBadRequest, "uploadresult.html", res)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		res.Result = "上传失败！"
		res.Desc = "can not read post file param."
		c.HTML(http.StatusBadRequest, "uploadresult.html", res)
		return
	}
	if file.Size > 50*1024*1024 {
		res.Result = "上传失败！"
		res.Desc = "can not post file large than " + strconv.Itoa(50*1024*1024)
		c.HTML(http.StatusInternalServerError, "uploadresult.html", res)
		return
	}

	newfilename := "videos/" + util.NewUUID() + ".mp4"
	err = c.SaveUploadedFile(file, newfilename)
	if err != nil {
		res.Result = "上传失败！"
		res.Desc = "save file failed."
		c.HTML(http.StatusInternalServerError, "uploadresult.html", res)
		return
	} else {
		video := &model.Video{
			Title:       title,
			Description: description,
			Filepath:    newfilename,
			Userid:      user,
		}
		if !dao.AddVideo(video) {
			res.Result = "上传失败！"
			res.Desc = "insert database failed."
			c.HTML(http.StatusInternalServerError, "uploadresult.html", res)
			return
		} else {
			res.Result = "上传成功！"
			res.Desc = "post succeed."
			c.HTML(http.StatusOK, "uploadresult.html", res)
			return
		}
	}
}
