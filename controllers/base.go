package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

type BaseController struct {
	beego.Controller
}

type response struct {
	ResponseCode int         `json:"response_code"`
	ResponseMsg  string      `json:"response_msg"`
	ResponseData interface{} `json:"response_data"`
}

func (this *BaseController) RenderJSON(code int, msg string, data interface{}) {
	resp := &response{code, msg, data}
	this.Data["json"] = resp
	this.ServeJSON()
}

func allowCORS() {
	beego.InsertFilter("*",
		beego.BeforeRouter,
		cors.Allow(&cors.Options{
			AllowAllOrigins: true,
			// AllowOrigins:     []string{"http://how"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
			AllowCredentials: true}))
}

func init() {
	allowCORS()
}
