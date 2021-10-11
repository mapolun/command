package core

import (
	"command/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

//gorm链接数据库
func DbInit() error {
	db, err := gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
			config.Mysql.Master.UserName,
			config.Mysql.Master.Password,
			config.Mysql.Master.Host,
			config.Mysql.Master.Port,
			config.Mysql.Master.Database,
		),
	)

	if err != nil {
		return err
	}
	/*连接池信息*/
	db.DB().SetMaxIdleConns(config.Mysql.Master.MaxIdleConns) //设置最大空闲数
	db.DB().SetMaxOpenConns(config.Mysql.Master.MaxOpenConns) //设置最大连接数
	db.SingularTable(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "ww_" + defaultTableName
	}
	Db = db
	return nil
}
