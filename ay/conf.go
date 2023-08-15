/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package ay

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

func InitConfig() {
	Yaml = getConfig()
}

func getConfig() *viper.Viper {
	config := viper.New()
	config.SetConfigName("config")
	config.AddConfigPath(".")
	config.SetConfigType("yaml")
	err := config.ReadInConfig()
	if err != nil {
		log.Printf("Failed to get the configuration.")
	}
	return config
}

func WatchConf() {
	Yaml.WatchConfig()
	Yaml.OnConfigChange(func(event fsnotify.Event) {
		// 配置文件修改重新执行的方法
		Init()
		Logger.Info("Detect config change: " + event.String())
	})
}
