package controllers

import (
	"fmt"
	"recitationSquare/snowflake"
	"strconv"
)

type MainController struct {
	BaseController
}

func (this *MainController) Get() {

	uniqueID, err := strconv.Atoi(snowflake.GetUniqueID())
	if err != nil {
		fmt.Println("转换失败")
	}
	this.renderJSON(uniqueID, "这是消息", nil)
	// this.renderSuccessJSON("成功", nil)
}
