package internal

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
func (c *BaseController) CheckParams(params interface{}) bool {
	if err := c.ParseForm(params); err != nil {
		/*
			指针错误一般不会发生，可能发生的错误在
			于解析时类型不匹配，即参数错误
		*/
		c.RenderParamsErrorJSON("参数错误", nil)
		return false
	}

	//验证参数
	valid := validation.Validation{}
	if ok, _ := valid.Valid(params); ok {
		return true
	} else {
		for _, err := range valid.Errors {
			c.RenderParamsErrorJSON(err.Message, nil)
			return false
		}
	}

	c.RenderParamsErrorJSON("参数错误", nil)
	return false
}

type ResponseData map[string]interface{}

func (c *BaseController) RenderJSON(code int, msg string, data interface{}, name ...string) {
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
		if len(name) > 0 {
			res[name[0]] = data
		} else {
			res["F_data"] = data
		}
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
	logger.Info("响应请求", requestInfo)
}

func (c *BaseController) RenderSuccessJSON(msg string, data interface{}, name ...string) {
	c.RenderJSON(model.SUCCESS_CODE, msg, data, name...)
}

// 错误响应
func (c *BaseController) RenderErrorJSON(err error, data interface{}, name ...string) {
	// 如果是定义的model error，有相应的code类型
	// 不会存在类型为ModelError但是值为nil的情况
	if mErr, ok := err.(*model.ModelError); ok {
		c.RenderJSON(mErr.Code, mErr.Msg, data, name...)
		return
	}

	// db error
	if dbErr, ok := err.(*model.DBError); ok {
		c.RenderJSON(model.DBERROR_CODE, dbErr.Msg, data, name...)
		return
	}

	c.RenderUnknownErrorJSON(err.Error(), data, name...)
}

// 参数错误  status 400
func (c *BaseController) RenderParamsErrorJSON(msg string, data interface{}, name ...string) {
	c.RenderJSON(model.PARAMS_ERROR_CODE, msg, data, name...)
}

// 未知错误，多用于内部处理错误或者不确定的错误
func (c *BaseController) RenderUnknownErrorJSON(msg string, data interface{}, name ...string) {
	c.RenderJSON(model.ERROR_CODE, msg, data, name...)
}
