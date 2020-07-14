package controllers

type MainController struct {
	BaseController
}

type Response struct {
	ResponseCode int    `json:"response_code"`
	ResponseMsg  string `json:"response_data"`
}

func (this *MainController) Get() {
	this.RenderJSON(10000, "成功", "没有数据")
}
