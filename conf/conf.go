/*
 * @Author: gongluck
 * @Date: 2020-06-17 14:43:27
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-17 15:52:31
 */

package conf

import (
	"fmt"

	"github.com/jinzhu/configor"
)

var Config = struct {
	APPName string `default:"govideo_server_debug"`
	Http    struct {
		// gin
		GinMode string `default:"debug"`
		// http
		HttpAddr    string `default:":80"`
		SessionName string `default:"govideo_server"`
	}
	Redis struct {
		ConnSize int    `default:"10"`
		NetWork  string `default:"tcp"`
		Address  string `default:"localhost:6379"`
		Password string `default:""`
		Key      string `default:"govideo_server"`
	}
	Video struct {
		MaxFileSize   int64  `default:"20971520"`
		FilePrefix    string `default:"./videos/"`
		TemplatesPath string `default:"./templates/"`
	}
}{}

func init() {
	configor.New(&configor.Config{Verbose: true}).Load(&Config, "./conf/config.yml")

	fmt.Println("config : *****************")
	fmt.Println("APPName :", Config.APPName)
	fmt.Println("Http :", Config.Http.GinMode, Config.Http.HttpAddr, Config.Http.SessionName)
	fmt.Println("Redis :", Config.Redis.Address, Config.Redis.ConnSize, Config.Redis.Key, Config.Redis.NetWork, Config.Redis.Password)
	fmt.Println("Videos :", Config.Video.FilePrefix, Config.Video.MaxFileSize, Config.Video.TemplatesPath)
}
