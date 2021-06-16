package controllers

import (
	"golang-project-prototype/library/util"
	"golang-project-prototype/library/util/logger"
	"golang-project-prototype/model"
	"net/url"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	values := url.Values{}
	values.Add("postId", "1")

	var resp interface{}
	err := util.GetHttpClient(model.CLIENT_DEFAULT).
		GetByReceiver("http://jsonplaceholder.typicode.com/comments", values, &resp)
	if err != nil {
		logger.Error("请求失败", err)
		c.renderErrorJSON(err, nil)
		return
	}

	c.renderSuccessJSON("成功", resp)

}
