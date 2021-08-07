// db migrate
package db

import (
	"database/sql"
	"golang-project-prototype/config"
	"golang-project-prototype/library/util/logger"

	"fmt"
	"path/filepath"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type MLog struct {
}

func (MLog) Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}
func (MLog) Verbose() bool {
	return true
}

func migration() {

	var err error
	dbHost := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True", config.DBConfig.DBUsername, config.DBConfig.DBPassword, config.DBConfig.DBHost, config.DBConfig.DBName)
	logger.Info("dbhost: ", dbHost)
	db, err := sql.Open("mysql", dbHost+"&loc=Asia%2FShanghai&allowNativePasswords=true&multiStatements=true")
	if err != nil {
		logger.Error("migration:初始化数据库失败", err)
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		logger.Error("migration:初始化db.Driver失败", err)
	}
	realPath, _ := filepath.Abs("migration")
	logger.Info("migration:filepath: ", realPath)

	realPath = strings.Replace(realPath, "\\", "/", -1)
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+realPath,
		"mysql",
		driver,
	)
	if err != nil {
		logger.Error("migration: 初始化migration失败", err)
	}
	var mLog MLog
	m.Log = mLog
	m.Up()
}
