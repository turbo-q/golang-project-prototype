package controllers

import (
	"demo/global"
	"net/http"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

type ResponseData map[string]interface{}

//JSON数据的返回

func (this *BaseController) renderJSON(code int, msg string, data interface{}) {
	resp := ResponseData{
		"F_responseNo":  code,
		"F_responseMsg": msg,
	}
	if data != nil {
		resp["F_data"] = data
	}

	if code == global.PARAMS_ERROR_CODE { //参数错误
		this.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
		this.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	}
	if code == global.TOKEN_ERROR_CODE { //token(access token , refresh access token) 错误
		this.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
		this.Ctx.ResponseWriter.WriteHeader(http.StatusForbidden)
	}

	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *BaseController) renderSuccessJSON(msg string, data interface{}) {
	this.renderJSON(global.SUCCESS_CODE, msg, data)
}

func (this *BaseController) renderErrorJSON(msg string, data interface{}) {
	this.renderJSON(global.ERROR_CODE, msg, data)
}

func (this *BaseController) renderParamsErrorJSON(msg string, data interface{}) {
	this.renderJSON(global.PARAMS_ERROR_CODE, msg, data)
}

// 没有权限
func (this *BaseController) echo403() {
	this.Ctx.ResponseWriter.WriteHeader(http.StatusForbidden)
	this.StopRun()
}

// 需要登录
func (this *BaseController) echo401() {
	this.Ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
	this.StopRun()
}

//重定向
func (this *BaseController) redirect(url string) {
	this.Redirect(url, http.StatusFound)
	this.StopRun()
}

//资源未修改
func (this *BaseController) echo304() {
	this.Ctx.ResponseWriter.WriteHeader(http.StatusNotModified)
	this.StopRun()
}
