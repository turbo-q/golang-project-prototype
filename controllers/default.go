package controllers

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.renderSuccess("成功", nil)
}
