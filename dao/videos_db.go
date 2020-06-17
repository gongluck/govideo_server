/*
 * @Author: gongluck
 * @Date: 2020-06-03 11:00:14
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-08 11:13:27
 */

package dao

import (
	"govideo_server/model"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func DelVideos() {
	db.Delete(&model.Video{})
}

func DelVideo(video *model.Video) {
	db.Delete(video)
}

func AddVideo(video *model.Video) bool {
	db.Create(&video)
	return !db.NewRecord(video)
}

func GetVideoByID(id int64) *model.Video {
	video := &model.Video{}
	db.First(video, id)
	return video
}

func GetVideoByTitle(title string) *model.Video {
	video := &model.Video{}
	db.Where("title=?", title).First(video)
	return video
}

func GetVideos() (videos []*model.Video) {
	db.Find(&videos)
	return videos
}

func GetVideosByLimit(limit, offset int) (videos []*model.Video) {
	db.Offset(offset).Limit(limit).Find(&videos)
	return videos
}

func GetVideosCount() (count int) {
	db.Model(&model.Video{}).Count(&count)
	return count
}
