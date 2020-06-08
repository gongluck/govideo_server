# govideo_server
go开发视频服务后台

## 数据库设计

### users表

|字段|描述|类型|
|:-|:-|-:|
|id|用户唯一标识(主键,自动增长)|INTEGER|
|name|用户名(非空,唯一)|text|
|password|密码(密文)|text|
|level|用户等级(非空)|INTEGER|

### videos表

|字段|描述|类型|
|:-|:-|-:|
|id|视频唯一标识(主键,自动增长)|INTEGER|
|title|视频标题(非空)|text|
|description|视频描述|text|
|filepath|视频文件路径(非空)|text|
|userid|所属用户id|INTEGER|

### 表结构定义
数据库模块使用[GORM](https://gorm.io/)
```Go
type User struct {
	//gorm.Model
	ID       uint   `gorm:"column:id;type:integer;primary_key;auto_increment"`
	Name     string `gorm:"column:name;type:text;not null;unique"`
	Password string `gorm:"column:password;type:text;not null"`
	Level    uint   `gorm:"column:level;type:integer"`
}
type Video struct {
	//gorm.Model
	ID          uint   `json:"id"gorm:"column:id;primary_key;auto_increment"`
	Title       string `json:"title"gorm:"column:title;type:text;not null"`
	Description string `json:"description"gorm:"column:description;type:text"`
	Filepath    string `json:"filepath"gorm:"column:filepath;not null"`
	Userid      uint   `json:"userid"gorm:"column:userid;type:integer;not null"`
}
```
