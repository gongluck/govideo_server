/*
 * @Author: gongluck
 * @Date: 2020-06-02 20:05:08
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-10 09:57:32
 */

package model

type User struct {
	//gorm.Model
	ID       int64  `gorm:"column:id;type:integer;primary_key;auto_increment"`
	Name     string `gorm:"column:name;type:text;not null;unique"`
	Password string `gorm:"column:password;type:text;not null"`
	Level    int64  `gorm:"column:level;type:integer"`
}
