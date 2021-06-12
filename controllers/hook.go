package controllers

import (
	"golang-project-prototype/library/util/logger"
	"net/http"
	"net/url"

	"github.com/astaxie/beego/context"
)

// 全局prepare方法，打印必要的请求信息
func GlobalPrepare(ctx *context.Context) {
	var (
		values url.Values
	)
	if ctx.Input.Method() == http.MethodGet {
		values = ctx.Request.URL.Query()
	} else {
		values = ctx.Request.Form
	}
	requestInfo := map[string]interface{}{
		"method":      ctx.Input.Method(),
		"request uri": ctx.Input.URI(),
		"values":      values.Encode(),
		"ip":          ctx.Input.IP(),
		"domain":      ctx.Input.Domain(),
	}
	logger.Infom("收到请求", requestInfo)
}
