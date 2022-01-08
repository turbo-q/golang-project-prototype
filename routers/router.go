package routers

import (
	"golang-project-prototype/config"
	v1 "golang-project-prototype/controllers/v1"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// cors config
func allowCORS() {
	beego.InsertFilter("*", beego.BeforeRouter, CORS)
}

// 全局的prepare钩子，打印一些必要的信息
func globalPrepare() {
	beego.InsertFilter("*", beego.BeforeRouter, GlobalPrepare)
}

func init() {
	allowCORS()
	globalPrepare()
	// router 使用namespace的形式进行构建
	// 结合restful能更好的表现层级关系
	ns :=
		beego.NewNamespace("/v1",
			//请求校验
			beego.NSCond(func(ctx *context.Context) bool {
				return ctx.Input.Query("apiToken") == config.DefaultConfig.ApiToken
			}),
			beego.NSRouter("/", &v1.MainController{}),
		)
	beego.AddNamespace(ns)
}
