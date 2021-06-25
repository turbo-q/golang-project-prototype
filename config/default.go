package config

import (
	"github.com/astaxie/beego/config"
)

var (
	DefaultConfig   defaultConfig
	DBConfig        dbConfig
	SnowflakeConfig snowflakeConfig
	RedisConfig     redisConfig
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

// redis config
type redisConfig struct {
	RedisMaxIdle     int
	RedisIdleTimeout int
	RedisHost        string
	RedisPassword    string
}

func init() {
	iniconf, _ := config.NewConfig("ini", "conf/app.conf")

	DefaultConfig.Env = iniconf.String("runmode")

	section := "golang-project-prototype"
	// default config
	DefaultConfig.HttpTimeout, _ = iniconf.Int(section + "::httpTimeout")
	DefaultConfig.ApiToken = iniconf.String(section + "::apiToken")

	// redis config
	RedisConfig.RedisIdleTimeout, _ = iniconf.Int(section + "::rdIdleTimeout")
	RedisConfig.RedisMaxIdle, _ = iniconf.Int(section + "::rdMaxIdle")
	RedisConfig.RedisPassword = iniconf.String(section + "::rdPassword")
	RedisConfig.RedisHost = iniconf.String(section + "::rdHost")

	// db config
	DBConfig.DBName = iniconf.String(section + "::dbName")
	DBConfig.DBUsername = iniconf.String(section + "::dbUsername")
	DBConfig.DBPassword = iniconf.String(section + "::dbPassword")
	DBConfig.DBHost = iniconf.String(section + "::dbHost")
	DBConfig.DBLog, _ = iniconf.Bool(section + "::dbLog")
	DBConfig.DBMaxIdle, _ = iniconf.Int(section + "::dbMaxIdle")
	DBConfig.DBMaxConn, _ = iniconf.Int(section + "::dbMaxConn")
	DBConfig.DBMaxConnLifetime, _ = iniconf.Int(section + "::dbMaxConnLifetime")

	// snowflake config
	section = "dreamSnowflake"
	SnowflakeConfig.Domain = iniconf.String(section + "::domain")
	SnowflakeConfig.AuthUser = iniconf.String(section + "::authUser")
	SnowflakeConfig.AuthSecret = iniconf.String(section + "::authUserSecurity")
}
