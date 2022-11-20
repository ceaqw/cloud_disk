package models

import (
	"CouldDisk/conf"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var MainDb *xorm.Engine

func init() {
	InitMainDb()
}

func InitMainDb() {
	mainDbConf := conf.GetDbCfgByName("maindb")
	jdbc := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		mainDbConf.UserName, mainDbConf.Password, mainDbConf.Host, mainDbConf.Port, mainDbConf.Database)
	MainDb, err := xorm.NewEngine(mainDbConf.Driver, jdbc)
	if err != nil {
		panic(fmt.Sprintf("maindb newEngine error: %#v\n", err.Error()))
	}
	err = MainDb.Ping()
	if err != nil {
		panic(fmt.Sprintf("maindb ping error: %#v\n", err.Error()))
	}
	MainDb.ShowSQL(true)
	err = MainDb.Sync(
		new(UserBasic),
	)
	if err != nil {
		panic(fmt.Sprintf("maindb sync error: %#v\n", err.Error()))
	}
}
