package db

import (
	"log"

	"github.com/astaxie/beego/config"
)

type dbConfig struct {
	dbName            string
	dbUsername        string
	dbPassword        string
	dbHost            string
	dbLog             bool
	dbMaxIdle         int
	dbMaxConn         int
	dbMaxConnLifetime int
}

var (
	DBConfig dbConfig
)

func dbConfigInit() {
	iniconf, err := config.NewConfig("ini", "conf/dev.conf")
	if err != nil {
		log.Println("加载conf/dev.conf失败")
	}

	section := "db-config"
	DBConfig = dbConfig{}
	DBConfig.dbName = iniconf.String(section + "::dbName")
	DBConfig.dbUsername = iniconf.String(section + "::dbUsername")
	DBConfig.dbPassword = iniconf.String(section + "::dbPassword")
	DBConfig.dbHost = iniconf.String(section + "::dbHost")
	DBConfig.dbLog, _ = iniconf.Bool(section + "::dbLog")
	DBConfig.dbMaxIdle, _ = iniconf.Int(section + "::dbMaxIdle")
	DBConfig.dbMaxConn, _ = iniconf.Int(section + "::dbMaxConn")
	DBConfig.dbMaxConnLifetime, _ = iniconf.Int(section + "::dbMaxConnLifetime")
}
