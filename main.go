package main

import (
	"command/autoload"
	"command/config"
	"command/core"
	"fmt"
	"github.com/ctfang/command"
)

func main() {
	// 加载配置
	config.Run()

	// 加载写入日志文件系统
	//logger.SetLogger("./config/logger.json")

	// 加载数据库
	if err := core.DbInit(); err != nil {
		fmt.Println("数据库链接失败", err)
		return
	}
	// 加载redis
	if err := core.RedisInit(); err != nil {
		fmt.Println("Redis链接失败", err)
		return
	}

	// 加载命令
	app := command.New()
	autoload.New(app)
	app.Run()
}
