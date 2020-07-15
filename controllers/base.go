package controllers

import (
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

func (this *BaseController) RenderJSON(code int, msg string, data interface{}) {
	resp := &response{code, msg, data}
	this.Data["json"] = resp
	this.ServeJSON()
}
