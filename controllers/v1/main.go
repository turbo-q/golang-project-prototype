package v1

import (
	"golang-project-prototype/library/util/http"
	"golang-project-prototype/library/util/logger"

	"golang-project-prototype/controllers/internal"
	"golang-project-prototype/model"
	"net/url"
)

type MainController struct {
	internal.BaseController
}

func (c *MainController) Get() {
	values := url.Values{}
	values.Add("postId", "1")

	var resp interface{}
	err := http.GetHttpClient(model.CLIENT_DEFAULT).
		GetByReceiver("http://jsonplaceholder.typicode.com/comments", values, &resp)
	if err != nil {
		logger.Error("请求失败", err)
		c.RenderErrorJSON(err, nil)
		return
	}

	c.RenderSuccessJSON("成功", resp)
}
