/*
 * @Author: gongluck
 * @Date: 2020-06-03 10:55:13
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-03 10:57:05
 */

package model

type Video struct {
	//gorm.Model
	ID          uint   `gorm:"column:id;primary_key;auto_increment"`
	Title       string `gorm:"column:title;type:text;not null"`
	Description string `gorm:"column:description;type:text"`
	Filepath    string `gorm:"column:filepath;not null"`
}
