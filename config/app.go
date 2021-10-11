package config

import (
	"fmt"
	"time"
)

//加载配置
func loading() {
	Mysql = mysqlConf{
		Master: mysqlModuleConf{
			Host:         "192.168.2.222",
			Database:     "xmt",
			UserName:     "root",
			Password:     "Ezsv9499",
			Port:         3306,
			MaxIdleConns: 20,
			MaxOpenConns: 300,
		},
	}
	Redis = redisConf{
		Master: redisModuleConf{
			Host:        "192.168.2.222",
			Password:    "",
			Port:        "6601",
			MaxIdle:     20,
			MaxActive:   300,
			IdleTimeout: 120,
			MaxTimeout:  time.Duration(30) * time.Second,
		},
	}
}

//初始化配置
func Run() {
	loading()
	fmt.Println("设置配置初始化成功")
}
