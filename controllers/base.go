package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type BaseController struct {
	beego.Controller
}

const (
	SUCCESS_CODE      = 10000 + iota
	ERROR_CODE        //未知错误
	PARAMS_ERROR_CODE // 参数错误
	TOKEN_ERROR_CODE  // token认证错误
	NOACCESS_CODE     // 未验证
)

// 获取服务端数据
type serverData struct {
	Data interface{} `json:"Data"`
	Msg  string
	OK   int `json:"OK"`
}

// 验证参数和解析数据
func (this *BaseController) checkParams(params interface{}) bool {
	if err := this.ParseForm(params); err != nil {
		/*
			指针错误一般不会发生，可能发生的错误在
			于解析时类型不匹配，即参数错误
		*/
		this.renderParamsErrorJSON("参数错误", nil)
		return false
	}

	//验证参数
	valid := validation.Validation{}
	if ok, _ := valid.Valid(params); ok {
		return true
	}

	this.renderParamsErrorJSON("参数错误", nil)
	return false
}

type ResponseData map[string]interface{}

func (this *BaseController) renderJSON(code int, msg string, data interface{}) {
	this.Ctx.Output.Header("Content-Type", "application/json")
	switch code {
	case PARAMS_ERROR_CODE:
		this.Ctx.Output.SetStatus(http.StatusBadRequest)
	case TOKEN_ERROR_CODE:
		this.Ctx.Output.SetStatus(http.StatusForbidden)
	case NOACCESS_CODE:
		this.Ctx.Output.SetStatus(http.StatusUnauthorized)
	case ERROR_CODE:
		this.Ctx.Output.SetStatus(http.StatusInternalServerError)
	default:
		this.Ctx.Output.SetStatus(http.StatusOK)
	}

	res := ResponseData{
		"F_responseNo":  code,
		"F_responseMsg": msg,
	}
	if data != nil {
		res["F_data"] = data
	}
	this.Data["json"] = res
	this.ServeJSON()
}

func (this *BaseController) renderSuccessJSON(msg string, data interface{}) {
	this.renderJSON(SUCCESS_CODE, msg, data)
}
func (this *BaseController) renderParamsErrorJSON(msg string, data interface{}) {
	this.renderJSON(PARAMS_ERROR_CODE, msg, data)
}

// 未知错误，多用于内部处理错误或者不确定的错误
func (this *BaseController) renderUnknownErrorJSON(msg string, data interface{}) {
	this.renderJSON(ERROR_CODE, msg, data)
}
