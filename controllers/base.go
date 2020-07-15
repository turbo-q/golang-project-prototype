package controllers

import (
	"demo/global"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

type response struct {
	ResponseCode int         `json:"response_code"`
	ResponseMsg  string      `json:"response_msg"`
	ResponseData interface{} `json:"response_data"`
}

func (this *BaseController) renderJSON(code int, msg string, data interface{}) {
	resp := &response{code, msg, data}
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *BaseController) renderSuccess(msg string, data interface{}) {
	this.renderJSON(global.SUCCESS_CODE, msg, data)
}

func (this *BaseController) renderError(msg string, data interface{}) {
	this.renderJSON(global.ERROR_CODE, msg, data)
}

func (this *BaseController) renderParamsError(msg string, data interface{}) {
	this.renderJSON(global.PARAMS_ERROR_CODE, msg, data)
}
