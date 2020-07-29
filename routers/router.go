package routers

import (
	"recitationSquare/controllers"

	"github.com/astaxie/beego"
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
	beego.Router("/", &controllers.MainController{})
}
