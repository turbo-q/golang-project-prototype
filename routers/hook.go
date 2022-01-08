package routers

import (
	"golang-project-prototype/library/util/logger"
	"net/http"
	"net/url"

	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
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
	logger.Info("收到请求", requestInfo)
}

// CORS cross-origin resource share
func CORS(ctx *context.Context) {
	cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		// AllowOrigins:     []string{"http://how"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true})
}
