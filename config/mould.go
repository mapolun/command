package config

import "time"

var (
	//Mysql 数据库信息
	Mysql mysqlConf
	//Redis 数据库信息
	Redis redisConf
	//TimeFormat 格式化时间
	TimeFormat = "2006-01-02 15:04:05"
	//Limit 分页一页显示数量
	Limit = 20
	//PathRoot 项目静态资源目录
	PathRoot = "/"
)

type mysqlConf struct {
	Master mysqlModuleConf
}

type redisConf struct {
	Master redisModuleConf
}

type mysqlModuleConf struct {
	Host         string
	Database     string
	UserName     string
	Password     string
	Port         int
	MaxIdleConns int /*用于设置闲置的连接数。*/
	MaxOpenConns int /*用于设置最大打开的连接数，默认值为0表示不限制*/
}

type redisModuleConf struct {
	Host        string
	Password    string
	Port        string
	MaxIdle     int           /*用于设置闲置的连接数。*/
	MaxActive   int           /*连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配*/
	IdleTimeout time.Duration /*连接关闭时间 300秒 （300秒不使用自动关闭）*/
	MaxTimeout  time.Duration
}

type codeConf struct {
	Ok int
	No int
}
