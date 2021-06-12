package controllers

import (
	"golang-project-prototype/model"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type BaseController struct {
	beego.Controller
}

// 验证参数和解析数据
func (c *BaseController) checkParams(params interface{}) bool {
	if err := c.ParseForm(params); err != nil {
		/*
			指针错误一般不会发生，可能发生的错误在
			于解析时类型不匹配，即参数错误
		*/
		c.renderParamsErrorJSON("参数错误", nil)
		return false
	}

	//验证参数
	valid := validation.Validation{}
	if ok, _ := valid.Valid(params); ok {
		return true
	} else {
		for _, err := range valid.Errors {
			c.renderParamsErrorJSON(err.Message, nil)
			return false
		}
	}

	c.renderParamsErrorJSON("参数错误", nil)
	return false
}

type ResponseData map[string]interface{}

func (c *BaseController) renderJSON(code int, msg string, data interface{}) {
	c.Ctx.Output.Header("Content-Type", "application/json")
	switch code {
	case model.PARAMS_ERROR_CODE:
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
	case model.TOKEN_ERROR_CODE:
		c.Ctx.Output.SetStatus(http.StatusForbidden)
	case model.NOACCESS_CODE:
		c.Ctx.Output.SetStatus(http.StatusUnauthorized)
	case model.ERROR_CODE:
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
	default:
		c.Ctx.Output.SetStatus(http.StatusOK)
	}

	res := ResponseData{
		"F_responseNo":  code,
		"F_responseMsg": msg,
	}
	if data != nil {
		res["F_data"] = data
	}
	c.Data["json"] = res
	c.ServeJSON()
}

func (c *BaseController) renderSuccessJSON(msg string, data interface{}) {
	c.renderJSON(model.SUCCESS_CODE, msg, data)
}

func (c *BaseController) renderErrorJSON(msg string, data interface{}) {
	c.renderJSON(model.ERROR_CODE, msg, data)
}

// 参数错误  status 400
func (c *BaseController) renderParamsErrorJSON(msg string, data interface{}) {
	c.renderJSON(model.PARAMS_ERROR_CODE, msg, data)
}

// 未知错误，多用于内部处理错误或者不确定的错误
func (c *BaseController) renderUnknownErrorJSON(msg string, data interface{}) {
	c.renderJSON(model.ERROR_CODE, msg, data)
}
