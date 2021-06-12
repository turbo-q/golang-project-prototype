package routers

import (
	"golang-project-prototype/config"
	"golang-project-prototype/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
)

func allowCORS() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		// AllowOrigins:     []string{"http://how"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true}))
}

// 全局的prepare钩子，打印一些必要的信息
func globalPrepare() {
	beego.InsertFilter("*", beego.BeforeRouter, controllers.GlobalPrepare)
}

func init() {
	allowCORS()
	globalPrepare()
	ns :=
		beego.NewNamespace("/api",
			//请求校验
			beego.NSCond(func(ctx *context.Context) bool {
				return ctx.Input.Query("apiToken") == config.DefaultConfig.ApiToken
			}),
		)
	beego.AddNamespace(ns)
}
