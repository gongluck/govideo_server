/*
 * @Author: gongluck
 * @Date: 2020-06-03 10:57:39
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-08 11:22:33
 */

package dao

import (
	"govideo_server/model"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// create database file, if it doesn't exist
	os.OpenFile("./database/govideo.db", os.O_RDONLY|os.O_CREATE, 0666)
	db, err = gorm.Open("sqlite3", "./database/govideo.db")
	if err != nil {
		panic("failed to connect database " + err.Error())
	}

	// 自动创建或更新表结构
	db.AutoMigrate(&model.User{}, &model.Video{})
}

func Close() error {
	return db.Close()
}
