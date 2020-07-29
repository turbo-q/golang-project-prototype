package controllers

import (
	"net/http"
	"recitationSquare/global/resp"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	A string
}

type ResponseData map[string]interface{}

//JSON数据的返回

func (this BaseController) Test() string {
	return "测试"
}

func (this *BaseController) renderJSON(code int, msg string, data interface{}) {
	res := ResponseData{
		"F_responseNo":  code,
		"F_responseMsg": msg,
	}
	if data != nil {
		res["F_data"] = data
	}

	if code == resp.PARAMS_ERROR_CODE { //参数错误
		this.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
		this.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	}
	if code == resp.TOKEN_ERROR_CODE { //token(access token , refresh access token) 错误
		this.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
		this.Ctx.ResponseWriter.WriteHeader(http.StatusForbidden)
	}

	this.Data["json"] = res
	this.ServeJSON()
}

func (this *BaseController) renderSuccessJSON(msg string, data interface{}) {
	this.renderJSON(resp.SUCCESS_CODE, msg, data)
}

func (this *BaseController) renderErrorJSON(msg string, data interface{}) {
	this.renderJSON(resp.ERROR_CODE, msg, data)
}

// 参数错误  status 400
func (this *BaseController) renderParamsErrorJSON(msg string, data interface{}) {
	this.renderJSON(resp.PARAMS_ERROR_CODE, msg, data)
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

//重定向302
func (this *BaseController) redirect(url string) {
	this.Redirect(url, http.StatusFound)
	this.StopRun()
}

//资源未修改
func (this *BaseController) echo304() {
	this.Ctx.ResponseWriter.WriteHeader(http.StatusNotModified)
	this.StopRun()
}
