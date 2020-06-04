/*
 * @Author: gongluck
 * @Date: 2020-06-04 09:48:42
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-04 10:01:04
 */
package handler

import (
	"govideo_server/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WebGetVideos(c *gin.Context) {
	videos := dao.GetVideos()
	c.HTML(http.StatusOK, "index.html", videos)
}
