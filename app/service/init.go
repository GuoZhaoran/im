package service

import (
	"IMtest/app/model"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DbEngine *xorm.Engine

func init() {
	drivername := "mysql"
	DsName := "root:root@(127.0.0.1:3306)/instant_messaging?"
	err := errors.New("")
	DbEngine, err = xorm.NewEngine(drivername, DsName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	//是否显示SQL语句
	DbEngine.ShowSQL(true)
	//数据库最大打开的连接数
	DbEngine.SetMaxOpenConns(10)

	//自动创建表结构
	DbEngine.Sync2(new(model.User), new(model.Community), new(model.Contact))
	fmt.Println("init data base ok")
}
