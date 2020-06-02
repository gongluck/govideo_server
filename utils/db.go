/*
 * @Author: gongluck
 * @Date: 2020-06-02 20:22:59
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-02 22:51:38
 */

package utils

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
}

func DelUsers() {
	db.Delete(&model.User{})
}

func AddUser(user *model.User) bool {
	db.Create(&user)
	return !db.NewRecord(user)
}
