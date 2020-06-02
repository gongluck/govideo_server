/*
 * @Author: gongluck
 * @Date: 2020-06-02 20:05:08
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-02 22:35:59
 */

package model

type User struct {
	//gorm.Model
	ID       uint   `gorm:"column:id;primary_key;auto_increment"`
	Name     string `gorm:"column:name;type:text;not null"`
	Password string `gorm:"column:password;type:text;not null"`
	Level    uint   `gorm:"column:level"`
}
