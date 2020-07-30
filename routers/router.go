package routers

import (
	"recitationSquare/global"

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

func init() {
	allowCORS()
	ns :=
		beego.NewNamespace("/api",
			//请求校验
			beego.NSCond(func(ctx *context.Context) bool {
				if ctx.Input.Query("apiToken") == global.API_TOKEN {
					return true
				}
				return false
			}),
		)
	beego.AddNamespace(ns)
	// beego.Router("/", &controllers.MainController{})
}
