# govideo_server
![](https://github.com/gongluck/govideo_server/blob/master/videos/logo.png)

[http://www.gongluck.icu/web/](http://www.gongluck.icu/web/)

[http://47.115.57.81/web/](http://47.115.57.81/web/)

## go开发视频后台服务
***GO***+***GORM***+***GIN***+***REDIS***+***UUID***+***DOCKER***

## ~~编译和运行~~(可以使用docker了)
```shell
go build
启动 redis
启动 govideo_server
浏览器访问 http://localhost/web
```

## docker编译运行
```shell
#下载更新镜像
docker pull gongluck/govideo_server

#手动运行
docker run -d --name redis redis
docker run -i -t --name govideo -p 80:80 -v /e/code/govideo_server/videos:/govideo_server/videos -v /e/code/govideo_server/conf:/govideo_server/conf -v /e/code/govideo_server/database:/govideo_server/database --link redis:redis gongluck/govideo_server

#脚本运行
wget https://raw.githubusercontent.com/gongluck/govideo_server/master/docker-compose.yml
mkdir conf
cd conf
wget https://raw.githubusercontent.com/gongluck/govideo_server/master/conf/config.yml
cd ..
docker-compose up -d
docker-compose down
```

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
	ID       int64   `gorm:"column:id;type:integer;primary_key;auto_increment"`
	Name     string `gorm:"column:name;type:text;not null;unique"`
	Password string `gorm:"column:password;type:text;not null"`
	Level    int64   `gorm:"column:level;type:integer"`
}
type Video struct {
	//gorm.Model
	ID          int64  `json:"id"gorm:"column:id;primary_key;auto_increment"`
	Title       string `json:"title"gorm:"column:title;type:text;not null"`
	Description string `json:"description"gorm:"column:description;type:text"`
	Filepath    string `json:"filepath"gorm:"column:filepath;not null"`
	Userid      int64  `json:"userid"gorm:"column:userid;type:integer;not null"`
}
```

## 会话管理

使用[Gin](https://gin-gonic.com/zh-cn/docs/)的**session**和其封装的**Redis**。

*windows*版本的*redis*使用这个版本[redis](https://github.com/microsoftarchive/redis/releases)
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

func SetSession(c *gin.Context, userid int64) error {
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

func GetUserID(c *gin.Context) int64 {
	session := sessions.Default(c)
	vsession := session.Get("session")
	user := session.Get(vsession)
	if user == nil {
		return 0
	} else {
		return user.(int64)
	}
}
```
