package config

import (
	"log"

	"github.com/astaxie/beego/config"
)

var (
	DBConfig dbConfig
)

type dbConfig struct {
	DBName            string
	DBUsername        string
	DBPassword        string
	DBHost            string
	DBLog             bool
	DBMaxIdle         int
	DBMaxConn         int
	DBMaxConnLifetime int
}

func init() {
	iniconf, err := config.NewConfig("ini", "conf/dev.conf")
	if err != nil {
		log.Println("加载conf/dev.conf失败")
	}

	section := "db-config"
	DBConfig = dbConfig{}
	DBConfig.DBName = iniconf.String(section + "::dbName")
	DBConfig.DBUsername = iniconf.String(section + "::dbUsername")
	DBConfig.DBPassword = iniconf.String(section + "::dbPassword")
	DBConfig.DBHost = iniconf.String(section + "::dbHost")
	DBConfig.DBLog, _ = iniconf.Bool(section + "::dbLog")
	DBConfig.DBMaxIdle, _ = iniconf.Int(section + "::dbMaxIdle")
	DBConfig.DBMaxConn, _ = iniconf.Int(section + "::dbMaxConn")
	DBConfig.DBMaxConnLifetime, _ = iniconf.Int(section + "::dbMaxConnLifetime")
}
