/*
 * @Author: gongluck
 * @Date: 2020-06-02 20:22:59
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-03 10:59:53
 */

package dao

import (
	"govideo_server/model"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func DelUsers() {
	db.Delete(&model.User{})
}

func AddUser(user *model.User) bool {
	db.Create(&user)
	return !db.NewRecord(user)
}

func GetUserByID(id uint) *model.User {
	user := &model.User{}
	db.First(user, id)
	return user
}

func GetUserByName(name string) *model.User {
	user := &model.User{}
	db.Where("name=?", name).First(user)
	return user
}

func GetUsers() (users []*model.User) {
	db.Find(&users)
	return users
}
