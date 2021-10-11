package cmd

import (
	"command/core"
	"command/model"
	"fmt"
	"github.com/ctfang/command"
	"github.com/gookit/color"
	"github.com/wonderivan/logger"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Foo struct{}

func (Foo) Configure() command.Configure {
	return command.Configure{
		Name:        "foo",
		Description: "示例脚本",
		Input: command.Argument{
			// 可选的参数，不输入也能执行
			Argument: []command.ArgParam{
				{Name: "id", Description: "媒体ID"},
			},
		},
	}
}

var (
	count int32
	m     sync.Mutex
)

func (Foo) Execute(input command.Input) {
	id, _ := strconv.Atoi(input.GetOption("id"))
	for {
		time.Sleep(time.Second * 2)
		//命令参数接收示例
		//redis操作示例
		redis := core.RedisConn.Get()
		_, err := redis.Do("Set", "loading", "123")
		if err != nil {
			fmt.Println(err)
			return
		}

		for i := 1; i <= 10; i++ {
			go selects(id)
		}
		break
	}
	time.Sleep(time.Second * 3)

	logger.Info(color.BgBlue.Sprintf("数量：%v", count))
}

func selects(id int) {
	m.Lock() //互斥
	value := count
	runtime.Gosched()
	value++
	count = value
	//atomic.AddInt32(&count,1) //安全写入
	m.Unlock()
	var data model.WxLogin
	//mysql操作示例
	core.Db.Table(model.WxLogin{}.TableName()).Where("id = ?", id).Scan(&data)

	fmt.Println(data)
}
