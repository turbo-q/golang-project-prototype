package routers

import (
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
			//此处正式版时改为验证加密请求
			beego.NSCond(func(ctx *context.Context) bool {
				if ctx.Input.Query("apiToken") == "CFIsGgvkonYEoVURomNZCk1HwshSQhDw" {
					return true
				}
				return false
			}),
		)
	beego.AddNamespace(ns)
	// beego.Router("/", &controllers.MainController{})
}
