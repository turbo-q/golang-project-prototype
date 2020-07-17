package controllers

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.renderSuccessJSON("成功", nil)
}
