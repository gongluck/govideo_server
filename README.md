# govideo_server
go开发视频服务后台

## 数据库设计

### users表

|字段|描述|类型|
|:-|:-|-:|
|id|用户唯一标识(主键,自动增长)|INTEGER|
|name|用户名(非空)|text|
|password|密码|text|
|level|用户等级(非空)|INTEGER|

### videos表

|字段|描述|类型|
|:-|:-|-:|
|id|视频唯一标识(主键,自动增长)|INTEGER|
|title|视频标题(非空)|text|
|description|视频描述|text|
|filepath|视频文件路径(非空)|text|

数据库模块使用[GORM](https://gorm.io/)
