package db

import (
	"fmt"
	_ "github.com/betterfor/BookDown/conf"
	"github.com/betterfor/GoLogger/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/spf13/viper"
)

var (
	engine *xorm.Engine
)

func init() {
	var err error
	user := viper.GetString("db.user")
	pwd := viper.GetString("db.passwd")
	ip := viper.GetString("db.ip")
	port := viper.GetInt("db.port")
	db := viper.GetString("db.db")
	showSql := viper.GetBool("db.showsql")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, pwd, ip, port, db) + "&loc=Asia%2FShanghai"
	fmt.Println(dataSource)
	engine, err = xorm.NewEngine("mysql", dataSource)
	if err != nil {
		panic(fmt.Sprintf("mysql init error: %v", err))
	}
	err = engine.DB().Ping()
	if err != nil {
		panic(fmt.Sprintf("mysql ping error: %v", err))
	}
	logger.Info("success connect mysql =====>", dataSource)
	engine.ShowSQL(showSql)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	engine.SetMaxIdleConns(5)  //设置连接池的空闲数大小
	engine.SetMaxOpenConns(30) //设置最大打开连接数
}

// 获取MySQL链接
func GetEngine() *xorm.Engine {
	return engine
}

// GetCheck 检查是否查询到,因为xorm不管真假
// 返回的结果为两个参数，
// 一个has为该条记录是否存在，第二个参数err为是否有错误。不管err是否为nil，has都有可能为true或者false。
func GetCheck(has bool, err error) error {
	if err != nil || !has {
		return fmt.Errorf("查询出错,可能的错误是:%v", err)
	}
	return nil
}
