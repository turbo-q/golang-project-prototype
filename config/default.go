package config

import (
	"github.com/astaxie/beego/config"
)

var (
	DefaultConfig   defaultConfig
	DBConfig        dbConfig
	SnowflakeConfig snowflakeConfig
)

// default config
type defaultConfig struct {
	HttpTimeout int
	ApiToken    string
	Env         string
}

// db config
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

// SnowflakeConfig 发号器配置
type snowflakeConfig struct {
	Domain     string
	AuthUser   string
	AuthSecret string
}

func init() {
	iniconf, _ := config.NewConfig("ini", "conf/app.conf")

	DefaultConfig.Env = iniconf.String("runmode")

	section := "golang-project-prototype"
	DefaultConfig.HttpTimeout, _ = iniconf.Int(section + "::httpTimeout")
	DefaultConfig.ApiToken = iniconf.String(section + "::apiToken")

	section = "db-config"
	DBConfig.DBName = iniconf.String(section + "::dbName")
	DBConfig.DBUsername = iniconf.String(section + "::dbUsername")
	DBConfig.DBPassword = iniconf.String(section + "::dbPassword")
	DBConfig.DBHost = iniconf.String(section + "::dbHost")
	DBConfig.DBLog, _ = iniconf.Bool(section + "::dbLog")
	DBConfig.DBMaxIdle, _ = iniconf.Int(section + "::dbMaxIdle")
	DBConfig.DBMaxConn, _ = iniconf.Int(section + "::dbMaxConn")
	DBConfig.DBMaxConnLifetime, _ = iniconf.Int(section + "::dbMaxConnLifetime")

	appName := "dreamSnowflake"
	SnowflakeConfig.Domain = iniconf.String(appName + "::domain")
	SnowflakeConfig.AuthUser = iniconf.String(appName + "::authUser")
	SnowflakeConfig.AuthSecret = iniconf.String(appName + "::authUserSecurity")
}
