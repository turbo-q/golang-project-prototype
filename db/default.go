package db

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dbDefault *gorm.DB

func init() {
	// 一些数据库初始化的配置信息config/DBConfig
	dbConfigInit()
	mysqlInit()
}

func mysqlInit() {

	db, err := gorm.Open("mysql", DBConfig.dbUsername+":"+DBConfig.dbPassword+"@("+DBConfig.dbHost+")/"+DBConfig.dbName+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Println("初始化数据库"+DBConfig.dbName+"失败：", err)
	}
	db.LogMode(DBConfig.dbLog)                                                          // 启用logger日志显示详细信息
	db.DB().SetMaxIdleConns(DBConfig.dbMaxIdle)                                         // 连接池最大闲置连接数量
	db.DB().SetMaxOpenConns(DBConfig.dbMaxConn)                                         // 数据库最大连接数量
	db.DB().SetConnMaxLifetime(time.Second * time.Duration(DBConfig.dbMaxConnLifetime)) // 最大可复用时间
	db.SingularTable(true)                                                              // 不使用复数表名

	dbDefault = db
}

func GetDBDefault() *gorm.DB {
	return dbDefault
}
