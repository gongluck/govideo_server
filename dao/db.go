/*
 * @Author: gongluck
 * @Date: 2020-06-03 10:57:39
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-03 10:59:19
 */

package dao

import (
	"fmt"
	"govideo_server/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	db, err = gorm.Open("sqlite3", "govideo.db")
	if err != nil {
		panic("failed to connect database")
	}

	// 检查模型`User`表是否存在
	if !db.HasTable(&model.User{}) {
		fmt.Println("CreateTable users")
		db.CreateTable(&model.User{})
	}

	if !db.HasTable(&model.Video{}) {
		fmt.Println("CreateTable videos")
		db.CreateTable(&model.Video{})
	}
}
