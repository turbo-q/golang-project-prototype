package snowflake

import (
	"github.com/astaxie/beego/config"
)

// SnowflakeConfig 发号器配置
type snowflakeConfig struct {
	Domain     string
	AuthUser   string
	AuthSecret string
}

var (
	SnowflakeConfig snowflakeConfig
)

func snowflakeConfigInit() {
	appName := "dreamSnowflake"
	appConf, _ := config.NewConfig("ini", "conf/app.conf")
	SnowflakeConfig = snowflakeConfig{}
	SnowflakeConfig.Domain = appConf.String(appName + "::domain")
	SnowflakeConfig.AuthUser = appConf.String(appName + "::authUser")
	SnowflakeConfig.AuthSecret = appConf.String(appName + "::authUserSecurity")
}
