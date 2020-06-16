/*
 * @Author: gongluck
 * @Date: 2020-06-03 10:55:13
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-10 15:37:03
 */

package model

type Video struct {
	//gorm.Model
	ID          int64  `json:"id"gorm:"column:id;primary_key;auto_increment"`
	Title       string `json:"title"gorm:"column:title;type:text;not null"`
	Description string `json:"description"gorm:"column:description;type:text"`
	Filepath    string `json:"filepath"gorm:"column:filepath;not null"`
	Userid      int64  `json:"userid"gorm:"column:userid;type:integer;not null"`
}
