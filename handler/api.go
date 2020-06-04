/*
 * @Author: gongluck
 * @Date: 2020-06-03 11:42:02
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-04 09:49:12
 */

package handler

import (
	"encoding/json"
	"fmt"
	"govideo_server/dao"
	"govideo_server/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReturnResult(c *gin.Context, status int, ret int, data interface{}) {
	c.JSON(status, gin.H{
		"ret":  ret,
		"data": data,
	})
}

func GetVideos(c *gin.Context) {
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
	file, err := c.FormFile("file")
	if err != nil {
		ReturnResult(c, http.StatusInternalServerError, -1, "can not read post file param.")
	}

	fmt.Println(file.Filename)
	err = c.SaveUploadedFile(file, "videos/"+file.Filename)
	if err != nil {
		ReturnResult(c, http.StatusInternalServerError, -1, "save file failed.")
	} else {
		title := c.PostForm("title")
		description := c.PostForm("description")
		video := &model.Video{
			Title:       title,
			Description: description,
			Filepath:    "videos/" + file.Filename,
		}
		if !dao.AddVideo(video) {
			ReturnResult(c, http.StatusInternalServerError, -1, "insert database failed.")
		} else {
			ReturnResult(c, http.StatusOK, 0, "post succeed.")
		}
	}
}
