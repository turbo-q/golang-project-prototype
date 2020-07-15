package controllers

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.echo403()
}
