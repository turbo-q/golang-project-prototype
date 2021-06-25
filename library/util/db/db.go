package db

import (
	"errors"
	"golang-project-prototype/config"
	"golang-project-prototype/library/util/logger"
	"golang-project-prototype/model"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dbDefault *gorm.DB

func init() {
	mysqlInit()
	migrate()
}

func mysqlInit() {
	dsn := config.DBConfig.DBUsername + ":" + config.DBConfig.DBPassword + "@(" + config.DBConfig.DBHost + ")/" + config.DBConfig.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", // 表名前缀，`User`表为`t_users`
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
	})
	if err != nil {
		log.Println("初始化数据库"+config.DBConfig.DBName+"失败：", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Error("获取sqlDB失败", err)
	}

	sqlDB.SetMaxIdleConns(config.DBConfig.DBMaxIdle)                                         // 连接池最大闲置连接数量
	sqlDB.SetMaxOpenConns(config.DBConfig.DBMaxConn)                                         // 数据库最大连接数量
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(config.DBConfig.DBMaxConnLifetime)) // 最大可复用时间                                                           // 不使用复数表名

	dbDefault = db
}

// db migrate
func migrate() {
	dbDefault.
		Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4").
		AutoMigrate(&model.User{})
}

func GetDBDefault() *gorm.DB {
	return dbDefault
}

// 由于gorm将not found也作业error
// 固进行区分严重错误
func GormErrorIsFatalError(d *gorm.DB) bool {
	return d.Error != nil && !GormErrorIsFatalError(d)
}

// not found,只有在first、last、take方法找不到时会ErrRecordNotFound
func GormErrorIsNotFound(d *gorm.DB) bool {
	return errors.Is(d.Error, gorm.ErrRecordNotFound)
}
