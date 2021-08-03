package controllers

import (
	"golang-project-prototype/library/util/logger"
	"golang-project-prototype/model"
	"net/http"
	"net/url"

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

// finish 打印输出
func (c *BaseController) Finish() {
	var (
		values url.Values
	)
	if c.Ctx.Input.Method() == http.MethodGet {
		values = c.Ctx.Request.URL.Query()
	} else {
		values = c.Ctx.Request.Form
	}
	requestInfo := map[string]interface{}{
		"method":      c.Ctx.Input.Method(),
		"request uri": c.Ctx.Input.URI(),
		"values":      values.Encode(),
		"content":     c.Data["json"],
	}
	logger.Infom("响应请求", requestInfo)
}

func (c *BaseController) renderSuccessJSON(msg string, data interface{}) {
	c.renderJSON(model.SUCCESS_CODE, msg, data)
}

// 错误响应
func (c *BaseController) renderErrorJSON(err error, data interface{}) {
	// 如果是定义的model error，有相应的code类型
	// 不会存在类型为ModelError但是值为nil的情况
	if mErr, ok := err.(*model.ModelError); ok {
		c.renderJSON(mErr.Code, mErr.Msg, data)
		return
	}

	// db error
	if dbErr, ok := err.(*model.DBError); ok {
		c.renderJSON(model.DBERROR_CODE, dbErr.Msg, data)
		return
	}

	c.renderUnknownErrorJSON(err.Error(), data)
}

// 参数错误  status 400
func (c *BaseController) renderParamsErrorJSON(msg string, data interface{}) {
	c.renderJSON(model.PARAMS_ERROR_CODE, msg, data)
}

// 未知错误，多用于内部处理错误或者不确定的错误
func (c *BaseController) renderUnknownErrorJSON(msg string, data interface{}) {
	c.renderJSON(model.ERROR_CODE, msg, data)
}
