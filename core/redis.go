package core

import (
	"command/config"
	"github.com/garyburd/redigo/redis"
)

var RedisConn *redis.Pool

//创建redis连接池
func RedisInit() error {
	pool := &redis.Pool{
		MaxIdle:     config.Redis.Master.MaxIdle,
		MaxActive:   config.Redis.Master.MaxActive,
		IdleTimeout: config.Redis.Master.IdleTimeout,
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			con, err := redis.Dial(
				"tcp",
				config.Redis.Master.Host+":"+config.Redis.Master.Port, // address
				redis.DialPassword(config.Redis.Master.Password),
				redis.DialConnectTimeout(config.Redis.Master.MaxTimeout),
				redis.DialReadTimeout(config.Redis.Master.MaxTimeout),
				redis.DialWriteTimeout(config.Redis.Master.MaxTimeout),
			)
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}
	RedisConn = pool
	return nil
}
