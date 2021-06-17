package db

import (
	"golang-project-prototype/config"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dbDefault *gorm.DB

func init() {
	mysqlInit()
}

func mysqlInit() {

	db, err := gorm.Open("mysql", config.DBConfig.DBUsername+":"+config.DBConfig.DBPassword+"@("+config.DBConfig.DBHost+")/"+config.DBConfig.DBName+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Println("初始化数据库"+config.DBConfig.DBName+"失败：", err)
	}
	db.LogMode(config.DBConfig.DBLog)                                                          // 启用logger日志显示详细信息
	db.DB().SetMaxIdleConns(config.DBConfig.DBMaxIdle)                                         // 连接池最大闲置连接数量
	db.DB().SetMaxOpenConns(config.DBConfig.DBMaxConn)                                         // 数据库最大连接数量
	db.DB().SetConnMaxLifetime(time.Second * time.Duration(config.DBConfig.DBMaxConnLifetime)) // 最大可复用时间
	db.SingularTable(true)                                                                     // 不使用复数表名

	dbDefault = db
}

func GetDBDefault() *gorm.DB {
	return dbDefault
}

// 由于gorm将not found也作业error
// 固进行区分严重错误
func GormErrorIsFatalError(d *gorm.DB) bool {
	return d.Error != nil && !GormErrorIsFatalError(d)
}

// not found
func GormErrorIsNotFound(d *gorm.DB) bool {
	return gorm.IsRecordNotFoundError(d.Error)
}
