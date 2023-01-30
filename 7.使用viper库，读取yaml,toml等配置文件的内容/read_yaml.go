package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// viper 灵活强大的配置管理工具，支持多种 JSON/TOML/YAML 等多种配置格式，支持热更新。
//https://darjun.github.io/2020/01/18/godailylib/viper/

// yaml file info -->

// database:
//  host: 127.0.0.1
//  user: root
//  dbname: test
//  pwd: 123456
func main() {
	//获取项目的执行路径
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	config := viper.New()
	config.AddConfigPath(path)   //设置读取的文件路径
	config.SetConfigName("d")    //设置读取的文件名
	config.SetConfigType("yaml") //设置文件的类型
	//尝试进行配置读取
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	//打印文件读取出来的内容:
	fmt.Println(config.Get("database.host"))
	fmt.Println(config.Get("database.user"))
	fmt.Println(config.Get("database.dbname"))
	fmt.Println(config.Get("database.pwd"))
}
