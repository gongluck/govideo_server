/*
 * @Author: gongluck
 * @Date: 2020-06-10 09:41:52
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-13 13:28:37
 */

package defs

var (
	// gin
	GinMode string = "debug"

	// http
	HttpAddr    string = ":80"
	SessionName string = "govideo_server"

	// redis
	RedisConnSize int    = 10
	RedisNetWork  string = "tcp"
	RedisAddress  string = "localhost:6379"
	RedisPassword string = ""
	RedisKey      string = "govideo_server"

	// videos
	MaxFileSize   int64  = 50 * 1024 * 1024
	FilePrefix    string = "./videos/"
	TemplatesPath string = "templates/"
)
