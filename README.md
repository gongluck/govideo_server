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

## 会话管理

使用[Gin](https://gin-gonic.com/zh-cn/docs/)的**session**和其封装的**Redis**。
```Go
package util

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func NewUUID() string {
	sessionid := uuid.NewV4()
	return sessionid.String()
}

func SetSession(c *gin.Context, userid uint) error {
	uuidstr := NewUUID()
	session := sessions.Default(c)
	session.Set("session", uuidstr)
	session.Set(uuidstr, userid)
	return session.Save()
}

func GetSessionAndUser(c *gin.Context) (interface{}, interface{}) {
	session := sessions.Default(c)
	vsession := session.Get("session")
	if vsession == nil {
		return nil, nil
	}

	user := session.Get(vsession)
	return vsession, user
}

func DelSession(c *gin.Context) error {
	session := sessions.Default(c)
	vsession, user := GetSessionAndUser(c)

	if vsession != nil {
		session.Delete(vsession)
	}
	if user != nil {
		session.Delete(user)
	}
	return session.Save()
}

func GetSessionUser(c *gin.Context) uint {
	session := sessions.Default(c)
	vsession := session.Get("session")
	user := session.Get(vsession)
	if user == nil {
		return 0
	} else {
		return user.(uint)
	}
}
```
